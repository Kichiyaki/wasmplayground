package assets

import "embed"

//go:embed *.js *.html *.wasm
var FS embed.FS
