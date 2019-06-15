package here

import (
	"encoding/json"
	"os"
	"strings"
)

type Info struct {
	Dir         string            `json:"Dir"`
	ImportPath  string            `json:"ImportPath"`
	Name        string            `json:"Name"`
	Doc         string            `json:"Doc"`
	Target      string            `json:"Target"`
	Root        string            `json:"Root"`
	Match       []string          `json:"Match"`
	Stale       bool              `json:"Stale"`
	StaleReason string            `json:"StaleReason"`
	GoFiles     []string          `json:"GoFiles"`
	Imports     []string          `json:"Imports"`
	Deps        []string          `json:"Deps"`
	TestGoFiles []string          `json:"TestGoFiles"`
	TestImports []string          `json:"TestImports"`
	Module      Module            `json:"Module"`
	GoEnv       map[string]string `json:"GoEnv"`
}

func (i Info) GoPath() string {
	return i.Getenv("GOPATH")
}

func (i Info) ModuleName() string {
	if i.Mods() {
		return i.Module.Path
	}
	return i.ImportPath
}

func (i Info) Getenv(k string) string {
	if i.GoEnv == nil {
		return ""
	}
	return i.GoEnv[k]
}

func (i Info) EnvHas(k string) bool {
	if i.GoEnv == nil {
		return false
	}
	_, ok := i.GoEnv[k]
	return ok
}

func (i Info) IsZero() bool {
	return i.String() == Info{}.String()
}

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
