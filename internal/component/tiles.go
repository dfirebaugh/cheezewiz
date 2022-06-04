package component

import (
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi"
)

type TileMapData struct {
	Map *tiled.Map
}

var TileMap = donburi.NewComponentType(TileMapData{})

func GetTileMap(entry *donburi.Entry) *TileMapData {
	return (*TileMapData)(entry.Component(TileMap))
}
