package entity

import (
	"cheezewiz/internal/component"

	"github.com/yohamta/donburi"
)

var DamageLabelTag = donburi.NewTag()

func MakeDamageLabel(w donburi.World, x float64, y float64, label string) {
	slabel := w.Create(DamageLabelTag, component.Position, component.ScreenLabel, component.Tick)

	entry := w.Entry(slabel)

	position := (*component.PositionData)(entry.Component(component.Position))
	l := (*component.LabelData)(entry.Component(component.ScreenLabel))
	ti := (*component.TickData)(entry.Component(component.Tick))

	ti.Value = 0
	ti.EOL = 100
	l.Label = label

	position.X = x
	position.Y = y

}
