package entity

import (
	"cheezewiz/assets"
	"cheezewiz/internal/component"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

var BackgroundTag = donburi.NewTag()

func MakeBackground(w donburi.World) {
	b := w.Create(BackgroundTag, component.Position, component.TileMap)
	entry := w.Entry(b)
	// position := (*component.PositionData)(entry.Component(component.Position))
	tiles := (*component.TileMapData)(entry.Component(component.TileMap))
	tm, err := assets.GetGardenMap()
	if err != nil {
		logrus.Panic(err.Error())
	}

	tiles.Map = tm
}