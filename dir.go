package here

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Dir attempts to gather info for the requested directory.
func (h Here) Dir(p string) (Info, error) {
	i, err := h.cache(p, func(p string) (Info, error) {
		var i Info

		fi, err := os.Stat(p)
		if err != nil {
			return i, err
		}

		if !fi.IsDir() {
			p = filepath.Dir(p)
		}

		pwd, err := os.Getwd()
		if err != nil {
			return i, err
		}

		defer os.Chdir(pwd)

		os.Chdir(p)

		b, err := run("go", "list", "-json")
		// go: cannot find main module; see 'go help modules'
		// build .: cannot find module for path .
		// no Go files in
		if err != nil {
			if nonGoDirRx.MatchString(err.Error()) {
				return fromNonGoDir(p)
			}
			return i, err
		}

		if err := json.Unmarshal(b, &i); err != nil {
			return i, err
		}

		return i, nil
	})

	if err != nil {
		return i, fmt.Errorf("here.Dir: %s: %w", p, err)
	}

	return h.cache(i.ImportPath, func(p string) (Info, error) {
		return i, nil
	})
}

// Dir attempts to gather info for the requested directory.
func Dir(p string) (Info, error) {
	return New().Dir(p)
}

func fromNonGoDir(dir string) (Info, error) {
	i := Info{
		Dir:  dir,
		Name: filepath.Base(dir),
	}

	b, err := run("go", "list", "-json", "-m")
	if err != nil {
		if nonGoDirRx.MatchString(err.Error()) {
			return i, nil
		}
		return i, fmt.Errorf("here.nonGoDir: %s: %w", dir, err)
	}

	// When workspaces used, go list returns multiple objects separated by new line
	// Solution is to treat is as an array
	b = append(append([]byte{'['}, b...), ']')

	rx := regexp.MustCompile("\\}(\r)?\n\\{")
	if rx.Match(b) {
		b = rx.ReplaceAll(b, []byte("},{"))
	}

	var Modules []Module
	if err := json.Unmarshal(b, &Modules); err != nil {
		return i, fmt.Errorf("here.nonGoDir: %s: %w", dir, err)
	}

	// Select current module by dir
	for _, module := range Modules {
		if strings.HasPrefix(dir, module.Dir) {
			i.Module = module
		}
	}

	if i.ImportPath == "" && i.Module.Path != "command-line-arguments" {
		i.ImportPath = i.Module.Path
	}

	return i, nil
}
