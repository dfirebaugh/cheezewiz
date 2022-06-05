package entity

import (
	"cheezewiz/internal/component"

	"github.com/yohamta/donburi"
)

var ExpBarTag = donburi.NewTag()

func MakeExpBar(w donburi.World) {
	expBar := w.Create(ExpBarTag, component.Exp, component.Position)
	entry := w.Entry(expBar)

	exp := (*component.ExpData)(entry.Component(component.Exp))

	*exp = component.ExpData{
		CurrentExp: 0,
		DesiredExp: 5, // initial
	}

	position := (*component.PositionData)(entry.Component(component.Position))
	*position = component.PositionData{
		X: 5,
		Y: 10,
	}
}