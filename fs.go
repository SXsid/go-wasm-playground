package root

import (
	"embed"
	"io/fs"
)

//go:embed assets/*

var embdedFs embed.FS

var AssetsFS fs.FS

// INFO: go run init funton in a packge when it get improrted
func init() {
	var err error
	AssetsFS, err = fs.Sub(embdedFs, "assets")
	if err != nil {
		panic(err)
	}
}
