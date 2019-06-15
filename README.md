# Here

[![GoDoc](https://godoc.org/github.com/gobuffalo/here?status.svg)](https://godoc.org/github.com/gobuffalo/here)

Here will get you **accurate** Go information about the directory of package requested.

## CLI

While you can use the tool via its API, you can also use the CLI to get a JSON version of the data.

### Installation

```bash
$ go get github.com/gobuffalo/here/cmd/here
```

### Usage

#### Default

```bash
$ here

{
  "Dir": "$GOPATH/src/github.com/gobuffalo/here",
  "ImportPath": "github.com/gobuffalo/here",
  "Name": "here",
  "Doc": "",
  "Target": "$GOPATH/pkg/darwin_amd64/github.com/gobuffalo/here.a",
  "Root": "$GOPATH",
  "Match": [
    "."
  ],
  "Stale": true,
  "StaleReason": "not installed but available in build cache",
  "GoFiles": [
    "here.go",
    "info.go",
    "module.go",
    "version.go"
  ],
  "Imports": [
    "encoding/json",
    "os",
    "os/exec",
    "path/filepath",
    "strings"
  ],
  "Deps": [
    "bytes",
    "context",
    "encoding",
    "encoding/base64",
    "encoding/binary",
    "encoding/json",
    "errors",
    "fmt",
    "internal/bytealg",
    "internal/cpu",
    "internal/fmtsort",
    "internal/poll",
    "internal/race",
    "internal/syscall/unix",
    "internal/testlog",
    "io",
    "math",
    "math/bits",
    "os",
    "os/exec",
    "path/filepath",
    "reflect",
    "runtime",
    "runtime/internal/atomic",
    "runtime/internal/math",
    "runtime/internal/sys",
    "sort",
    "strconv",
    "strings",
    "sync",
    "sync/atomic",
    "syscall",
    "time",
    "unicode",
    "unicode/utf16",
    "unicode/utf8",
    "unsafe"
  ],
  "TestGoFiles": null,
  "TestImports": null,
  "Module": {
    "Path": "github.com/gobuffalo/here",
    "Main": true,
    "Dir": "$GOPATH/src/github.com/gobuffalo/here",
    "GoMod": "$GOPATH/src/github.com/gobuffalo/here/go.mod",
    "GoVersion": "1.12"
  },
  "GoEnv": {
    "CC": "clang",
    "CGO_CFLAGS": "-g -O2",
    "CGO_CPPFLAGS": "",
    "CGO_CXXFLAGS": "-g -O2",
    "CGO_ENABLED": "1",
    "CGO_FFLAGS": "-g -O2",
    "CGO_LDFLAGS": "-g -O2",
    "CXX": "clang++",
    "GCCGO": "gccgo",
    "GOARCH": "amd64",
    "GOBIN": "",
    "GOCACHE": "$HOME/Library/Caches/go-build",
    "GOEXE": "",
    "GOFLAGS": "",
    "GOGCCFLAGS": "-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/zj/ktv0trrj4l79dfq0dkm1b6d40000gn/T/go-build762990333=/tmp/go-build -gno-record-gcc-switches -fno-common",
    "GOHOSTARCH": "amd64",
    "GOHOSTOS": "darwin",
    "GOMOD": "$GOPATH/src/github.com/gobuffalo/here/go.mod",
    "GOOS": "darwin",
    "GOPATH": "$GOPATH",
    "GOPROXY": "",
    "GORACE": "",
    "GOROOT": "/usr/local/go",
    "GOTMPDIR": "",
    "GOTOOLDIR": "/usr/local/go/pkg/tool/darwin_amd64",
    "PKG_CONFIG": "pkg-config"
  }
}
```

#### By Directory

```bash
$ here cmd/here

{
  "Dir": "$GOPATH/src/github.com/gobuffalo/here/cmd/here",
  "ImportPath": "github.com/gobuffalo/here/cmd/here",
  "Name": "main",
  "Doc": "",
  "Target": "$GOPATH/bin/here",
  "Root": "$GOPATH",
  "Match": [
    "."
  ],
  "Stale": false,
  "StaleReason": "",
  "GoFiles": [
    "main.go"
  ],
  "Imports": [
    "fmt",
    "github.com/gobuffalo/here",
    "log",
    "os"
  ],
  "Deps": [
    "bytes",
    "context",
    "encoding",
    "encoding/base64",
    "encoding/binary",
    "encoding/json",
    "errors",
    "fmt",
    "github.com/gobuffalo/here",
    "internal/bytealg",
    "internal/cpu",
    "internal/fmtsort",
    "internal/poll",
    "internal/race",
    "internal/syscall/unix",
    "internal/testlog",
    "io",
    "log",
    "math",
    "math/bits",
    "os",
    "os/exec",
    "path/filepath",
    "reflect",
    "runtime",
    "runtime/internal/atomic",
    "runtime/internal/math",
    "runtime/internal/sys",
    "sort",
    "strconv",
    "strings",
    "sync",
    "sync/atomic",
    "syscall",
    "time",
    "unicode",
    "unicode/utf16",
    "unicode/utf8",
    "unsafe"
  ],
  "TestGoFiles": null,
  "TestImports": null,
  "Module": {
    "Path": "github.com/gobuffalo/here",
    "Main": true,
    "Dir": "$GOPATH/src/github.com/gobuffalo/here",
    "GoMod": "$GOPATH/src/github.com/gobuffalo/here/go.mod",
    "GoVersion": "1.12"
  },
  "GoEnv": {
    "CC": "clang",
    "CGO_CFLAGS": "-g -O2",
    "CGO_CPPFLAGS": "",
    "CGO_CXXFLAGS": "-g -O2",
    "CGO_ENABLED": "1",
    "CGO_FFLAGS": "-g -O2",
    "CGO_LDFLAGS": "-g -O2",
    "CXX": "clang++",
    "GCCGO": "gccgo",
    "GOARCH": "amd64",
    "GOBIN": "",
    "GOCACHE": "$HOME/Library/Caches/go-build",
    "GOEXE": "",
    "GOFLAGS": "",
    "GOGCCFLAGS": "-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/zj/ktv0trrj4l79dfq0dkm1b6d40000gn/T/go-build976603462=/tmp/go-build -gno-record-gcc-switches -fno-common",
    "GOHOSTARCH": "amd64",
    "GOHOSTOS": "darwin",
    "GOMOD": "$GOPATH/src/github.com/gobuffalo/here/go.mod",
    "GOOS": "darwin",
    "GOPATH": "$GOPATH",
    "GOPROXY": "",
    "GORACE": "",
    "GOROOT": "/usr/local/go",
    "GOTMPDIR": "",
    "GOTOOLDIR": "/usr/local/go/pkg/tool/darwin_amd64",
    "PKG_CONFIG": "pkg-config"
  }
}
```

#### By Package

```bash
$ here pkg github.com/gobuffalo/genny

{
  "Dir": "$GOPATH/pkg/mod/github.com/gobuffalo/genny@v0.2.0",
  "ImportPath": "github.com/gobuffalo/genny",
  "Name": "genny",
  "Doc": "Package genny is a _framework_ for writing modular generators, it however, doesn't actually generate anything.",
  "Target": "",
  "Root": "",
  "Match": [
    "github.com/gobuffalo/genny"
  ],
  "Stale": true,
  "StaleReason": "build ID mismatch",
  "GoFiles": [
    "confirm.go",
    "dir.go",
    "disk.go",
    "dry_runner.go",
    "events.go",
    "file.go",
    "force.go",
    "generator.go",
    "genny.go",
    "group.go",
    "helpers.go",
    "logger.go",
    "replacer.go",
    "results.go",
    "runner.go",
    "step.go",
    "transformer.go",
    "version.go",
    "wet_runner.go"
  ],
  "Imports": null,
  "Deps": null,
  "TestGoFiles": [
    "dry_runner_test.go",
    "file_test.go",
    "force_test.go",
    "generator_test.go",
    "genny_test.go",
    "group_test.go",
    "helpers_test.go",
    "replacer_test.go",
    "results_test.go",
    "runner_test.go",
    "step_test.go",
    "transformer_test.go",
    "wet_runner_test.go"
  ],
  "TestImports": null,
  "Module": {
    "Path": "github.com/gobuffalo/genny",
    "Main": false,
    "Dir": "$GOPATH/pkg/mod/github.com/gobuffalo/genny@v0.2.0",
    "GoMod": "$GOPATH/pkg/mod/cache/download/github.com/gobuffalo/genny/@v/v0.2.0.mod",
    "GoVersion": ""
  },
  "GoEnv": {
    "CC": "clang",
    "CGO_CFLAGS": "-g -O2",
    "CGO_CPPFLAGS": "",
    "CGO_CXXFLAGS": "-g -O2",
    "CGO_ENABLED": "1",
    "CGO_FFLAGS": "-g -O2",
    "CGO_LDFLAGS": "-g -O2",
    "CXX": "clang++",
    "GCCGO": "gccgo",
    "GOARCH": "amd64",
    "GOBIN": "",
    "GOCACHE": "$HOME/Library/Caches/go-build",
    "GOEXE": "",
    "GOFLAGS": "",
    "GOGCCFLAGS": "-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/zj/ktv0trrj4l79dfq0dkm1b6d40000gn/T/go-build207410055=/tmp/go-build -gno-record-gcc-switches -fno-common",
    "GOHOSTARCH": "amd64",
    "GOHOSTOS": "darwin",
    "GOMOD": "$GOPATH/src/github.com/gobuffalo/here/go.mod",
    "GOOS": "darwin",
    "GOPATH": "$GOPATH",
    "GOPROXY": "",
    "GORACE": "",
    "GOROOT": "/usr/local/go",
    "GOTMPDIR": "",
    "GOTOOLDIR": "/usr/local/go/pkg/tool/darwin_amd64",
    "PKG_CONFIG": "pkg-config"
  }
}
```
