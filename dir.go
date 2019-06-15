package here

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var dirCache = &infoMap{}

// Dir attempts to gather info for the requested directory.
func Dir(p string) (Info, error) {
	var err error
	i, ok := dirCache.LoadOr(p, func(m *infoMap) (Info, bool) {
		i := newInfo()

		fi, err := os.Stat(p)
		if err != nil {
			return i, false
		}

		if !fi.IsDir() {
			p = filepath.Dir(p)
		}

		pwd, err := os.Getwd()
		if err != nil {
			return i, false
		}

		defer os.Chdir(pwd)

		os.Chdir(p)

		b, err := run("go", "list", "-json")
		if err != nil {
			return i, false
		}

		if err := json.Unmarshal(b, &i); err != nil {
			return i, false
		}

		if err := setEnv(&i); err != nil {
			return i, false
		}
		return i, true
	})

	if err != nil {
		return i, err
	}

	if !ok {
		return i, fmt.Errorf("an error occurred %s", p)
	}

	return i, nil
}
