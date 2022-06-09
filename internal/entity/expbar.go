package entity

import (
	"github.com/yohamta/donburi"
)

var ExpBarTag = donburi.NewTag()

func MakeExpBar(w donburi.World) {
	// expBar := w.Create(ExpBarTag, component.Exp, component.NextLevel)
	// entry := w.Entry(expBar)

	// exp := (*component.ExpData)(entry.Component(component.Exp))

	// *exp = component.ExpData{
	// 	CurrentExp: 0,
	// 	DesiredExp: 5, // initial
	// }

	// // position := (*component.PositionData)(entry.Component(component.Position))
	// // position.Set(5, 10)

	// nextLevel := (*component.NextLevelData)(entry.Component(component.NextLevel))

	// nextLevel.CurrentLevel = 1
}
