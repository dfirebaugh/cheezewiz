package assets

import (
	"bytes"
	"cheezewiz/assets/scenes/garden"
	"cheezewiz/assets/scenes/kitchen"
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
func GetKitchenMap() (*tiled.Map, error) {
	return tiled.LoadReader("./assets/scenes/kitchen", bytes.NewReader(kitchen.Kitchen1Raw))
}
