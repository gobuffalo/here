package here

import (
	"encoding/json"
	"os"
	"strings"
)

// Info represents details about the directory/package
type Info struct {
	Dir         string
	ImportPath  string
	Name        string
	Doc         string
	Target      string
	Root        string
	Match       []string
	Stale       bool
	StaleReason string
	GoFiles     []string
	Imports     []string
	Deps        []string
	TestGoFiles []string
	TestImports []string
	Module      Module
	GoEnv       map[string]string // go env -json
}

// GoPath returns the GOPATH ENV var
func (i Info) GoPath() string {
	return i.Getenv("GOPATH")
}

// ModuleName returns the name of the current
// module, or if not using modules, the current
// package. These *might* not match.
func (i Info) ModuleName() string {
	if i.Mods() {
		return i.Module.Path
	}
	return i.ImportPath
}

// Getenv is a helper function to mimic
// os.Getenv behavior.
func (i Info) Getenv(k string) string {
	if i.GoEnv == nil {
		return ""
	}
	return i.GoEnv[k]
}

// Has checks if the ENV variable is in the GoEnv
func (i Info) Has(k string) bool {
	if i.GoEnv == nil {
		return false
	}
	_, ok := i.GoEnv[k]
	return ok
}

// IsZero checks if the type has been filled
// with rich chocolately data goodness
func (i Info) IsZero() bool {
	return i.String() == Info{}.String()
}

// Mods returns whether Go modules are used
// in this directory/package.
func (i Info) Mods() bool {
	return !i.Module.IsZero()
}

func (i Info) String() string {
	if i.GoEnv == nil {
		return ""
	}
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err.Error()
	}
	s := string(b)
	s = strings.ReplaceAll(s, os.Getenv("GOPATH"), "$GOPATH")
	s = strings.ReplaceAll(s, os.Getenv("HOME"), "$HOME")
	return s
}
