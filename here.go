package here

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

func New() Info {
	return Info{
		GoEnv: map[string]string{},
	}
}

func Dir(p string) (Info, error) {
	i := New()

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
	if err != nil {
		return i, err
	}

	if err := json.Unmarshal(b, &i); err != nil {
		return i, err
	}

	if err := setEnv(&i); err != nil {
		return i, err
	}
	return i, nil
}

func Package(p string) (Info, error) {
	i := New()
	b, err := run("go", "list", "-json", "-find", p)
	if err != nil {
		return i, err
	}
	if err := json.Unmarshal(b, &i); err != nil {
		return i, err
	}

	if err := setEnv(&i); err != nil {
		return i, err
	}
	return i, nil
}

func run(n string, args ...string) ([]byte, error) {
	c := exec.Command(n, args...)

	b, err := c.Output()
	if err != nil {
		return b, err
	}

	return b, err
}

func setEnv(i *Info) error {
	i.GoEnv = map[string]string{}
	b, err := run("go", "env", "-json")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &i.GoEnv)

}
