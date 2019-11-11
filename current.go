package here

import (
	"path/filepath"
)

func (h Here) Current() (Info, error) {
	hp := &h
	(&hp.curOnce).Do(func() {
		b, err := run("go", "env", "GOMOD")
		if err != nil {
			hp.curErr = err
			return
		}
		root := filepath.Dir(string(b))
		i, err := h.Dir(root)
		if err != nil {
			hp.curErr = err
			return
		}
		hp.current = i
	})

	return h.current, h.curErr
}
