//go:build release

package static

import (
	"embed"
	"io/fs"
)

//go:embed public
var public embed.FS

var Public, _ = fs.Sub(public, "public")
