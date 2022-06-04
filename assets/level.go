package assets

import (
	"bytes"
	"cheezewiz/assets/scenes/garden"

	"github.com/lafriks/go-tiled"
)

func GetGardenMap() (*tiled.Map, error) {
	return tiled.LoadReader("", bytes.NewReader(garden.GardenRaw))
}
