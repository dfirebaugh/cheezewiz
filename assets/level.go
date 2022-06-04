package assets

import (
	"bytes"
	"cheezewiz/assets/scenes/garden"
	"cheezewiz/assets/scenes/kitchen"

	"github.com/lafriks/go-tiled"
)

func GetGardenMap() (*tiled.Map, error) {
	return tiled.LoadReader("", bytes.NewReader(garden.GardenRaw))
}

func GetKitchenMap() (*tiled.Map, error) {
	return tiled.LoadReader("", bytes.NewReader(kitchen.Kitchen1Raw))
}
