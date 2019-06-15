package here

import (
	"encoding/json"
	"os/exec"
)

func newInfo() Info {
	return Info{
		GoEnv: map[string]string{},
	}
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
