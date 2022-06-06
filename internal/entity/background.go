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
	tiles := (*component.TileMapData)(entry.Component(component.TileMap))
	tm, err := assets.GetKitchenMap()
	if err != nil {
		logrus.Panic(err)
	}

	tiles.Map = tm
}
