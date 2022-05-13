package here

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Dir(t *testing.T) {
	r := require.New(t)

	root, err := os.Getwd()
	r.NoError(err)

	info, err := Dir(root)
	r.NoError(err)
	sanityCheck(t, info)
}

func Test_Dir_NonGoDir(t *testing.T) {
	r := require.New(t)

	root, err := os.Getwd()
	r.NoError(err)

	empty := filepath.Join(root, "..", "empty")
	os.MkdirAll(empty, 0755)
	defer os.RemoveAll(empty)

	info, err := Dir(empty)
	r.NoError(err)
	r.NotZero(info)

	r.Equal(empty, info.Dir)
	r.Equal("", info.Module.GoMod)
	r.Equal("", info.ImportPath)
	r.Equal("empty", info.Name)
	r.Empty(info.GoFiles)
	r.Empty(info.Imports)
}

func Test_Dir_Within(t *testing.T) {
	r := require.New(t)

	root, err := os.Getwd()
	r.NoError(err)

	cmd := filepath.Join(root, "cmd")
	info, err := Dir(cmd)
	r.NoError(err)
	r.NotZero(info)

	r.Equal(cmd, info.Dir)
	r.Equal(filepath.Join(root, "go.mod"), info.Module.GoMod)
	r.Equal("github.com/gobuffalo/here", info.ImportPath)
	r.Equal("cmd", info.Name)
	r.Empty(info.GoFiles)
	r.Empty(info.Imports)
}

func Test_Dir_CmdHere(t *testing.T) {
	r := require.New(t)

	root, err := os.Getwd()
	r.NoError(err)

	cmd := filepath.Join(root, "cmd", "here")
	info, err := Dir(cmd)
	r.NoError(err)
	r.NotZero(info)

	r.Equal(cmd, info.Dir)
	r.Equal(filepath.Join(root, "go.mod"), info.Module.GoMod)
	r.Equal("github.com/gobuffalo/here/cmd/here", info.ImportPath)
	r.Equal("main", info.Name)
	r.NotEmpty(info.GoFiles)
	r.NotEmpty(info.Imports)
}

func Test_Dir_File(t *testing.T) {
	r := require.New(t)

	root, err := os.Getwd()
	r.NoError(err)

	cmd := filepath.Join(root, "cmd", "here", "main.go")
	info, err := Dir(cmd)
	r.NoError(err)
	r.NotZero(info)

	r.Equal(filepath.Dir(cmd), info.Dir)
	r.Equal(filepath.Join(root, "go.mod"), info.Module.GoMod)
	r.Equal("github.com/gobuffalo/here/cmd/here", info.ImportPath)
	r.Equal("main", info.Name)
	r.NotEmpty(info.GoFiles)
	r.NotEmpty(info.Imports)
}

func Test_Dir_Unknown(t *testing.T) {
	r := require.New(t)

	info, err := Dir("/unknown")
	r.Error(err)
	r.Zero(info)
}
