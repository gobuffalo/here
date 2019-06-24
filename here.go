package here

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
)

func newInfo() Info {
	return Info{
		GoEnv: map[string]string{},
	}
}

func run(n string, args ...string) ([]byte, error) {
	c := exec.Command(n, args...)

	bb := &bytes.Buffer{}
	c.Stdout = bb
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}

func setEnv(i *Info) error {
	i.GoEnv = map[string]string{}
	b, err := run("go", "env", "-json")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &i.GoEnv)

}
