
```bash
$ go list -json
```

```json
{
	"Dir": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/src/github.com/gobuffalo/here",
	"ImportPath": "github.com/gobuffalo/here",
	"Name": "here",
	"Target": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/pkg/darwin_amd64/github.com/gobuffalo/here.a",
	"Root": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go",
	"Module": {
		"Path": "github.com/gobuffalo/here",
		"Main": true,
		"Dir": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/src/github.com/gobuffalo/here",
		"GoMod": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/src/github.com/gobuffalo/here/go.mod",
		"GoVersion": "1.12"
	},
	"Match": [
		"."
	],
	"Stale": true,
	"StaleReason": "not installed but available in build cache",
	"GoFiles": [
		"env.go",
		"here.go",
		"info.go",
		"module.go",
		"version.go"
	],
	"Imports": [
		"bufio",
		"bytes",
		"encoding/json",
		"github.com/gobuffalo/here/internal/github.com/gobuffalo/syncx",
		"io",
		"os",
		"os/exec",
		"path/filepath",
		"strings"
	],
	"Deps": [
		"bufio",
		"bytes",
		"context",
		"encoding",
		"encoding/base64",
		"encoding/binary",
		"encoding/json",
		"errors",
		"fmt",
		"github.com/gobuffalo/here/internal/github.com/gobuffalo/syncx",
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
	]
}

```


```bash
$ go list -json -find github.com/gobuffalo/buffalo/render
```

```json
{
  "Dir": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/pkg/mod/github.com/gobuffalo/buffalo@v0.14.6/render",
  "ImportPath": "github.com/gobuffalo/buffalo/render",
  "Name": "render",
  "Module": {
    "Path": "github.com/gobuffalo/buffalo",
    "Version": "v0.14.6",
    "Time": "2019-06-14T14:10:43Z",
    "Indirect": true,
    "Dir": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/pkg/mod/github.com/gobuffalo/buffalo@v0.14.6",
    "GoMod": "/Users/markbates/Library/Mobile Documents/com~apple~CloudDocs/go/pkg/mod/cache/download/github.com/gobuffalo/buffalo/@v/v0.14.6.mod",
    "GoVersion": "1.12"
  },
  "Match": [
    "github.com/gobuffalo/buffalo/render"
  ],
  "Stale": true,
  "StaleReason": "build ID mismatch",
  "GoFiles": [
    "auto.go",
    "download.go",
    "func.go",
    "helpers.go",
    "html.go",
    "js.go",
    "json.go",
    "options.go",
    "plain.go",
    "render.go",
    "renderer.go",
    "sse.go",
    "string.go",
    "template.go",
    "template_engine.go",
    "template_helpers.go",
    "xml.go"
  ],
  "TestGoFiles": [
    "js_test.go",
    "partials_test.go",
    "render_test.go",
    "template_helpers_test.go"
  ],
  "XTestGoFiles": [
    "auto_test.go",
    "download_test.go",
    "func_test.go",
    "html_test.go",
    "json_test.go",
    "markdown_test.go",
    "plain_test.go",
    "string_test.go",
    "template_test.go",
    "xml_test.go"
  ]
}
```
