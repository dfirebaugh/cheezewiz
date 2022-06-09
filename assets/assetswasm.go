//go:build js
// +build js

package assets

import (
	"bytes"
	"cheezewiz/assets/scenes/garden"
	"embed"
	_ "embed"

	"github.com/lafriks/go-tiled"
)

//go:embed scenes/*
var SceneFS embed.FS

//go:embed *
var AssetFS embed.FS

func GetGardenMap() (*tiled.Map, error) {
	return tiled.LoadReader("./assets/scenes/garden", bytes.NewReader(garden.GardenRaw))
}

// we need to embed the map assets.  However, there is currently an issue preventing us from doing that
//   see: https://github.com/lafriks/go-tiled/issues/63
func GetKitchenMap() (*tiled.Map, error) {
	return &tiled.Map{}, nil
}
