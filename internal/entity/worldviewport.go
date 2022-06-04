package entity

import (
	"cheezewiz/internal/component"

	"github.com/yohamta/donburi"
)

var WorldViewPortTag = donburi.NewTag()

func MakeWorld(w donburi.World) {
	worldViewPort := w.Create(WorldViewPortTag, component.Position)
	w.Entry(worldViewPort)
}
