package here

import "path/filepath"

var root = func() string {
	b, _ := run("go", "env", "GOMOD")
	return filepath.Dir(string(b))
}()

func Current() (Info, error) {
	return Dir(root)
}
