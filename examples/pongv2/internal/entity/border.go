package entity

import (
	"cheezewiz/examples/pongv2/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

func NewTopBorder(w donburi.World) {
	screenWidth, _ := ebiten.WindowSize()
	tb := w.Create(component.Rect, component.Position, component.IsBat)
	entry := w.Entry(tb)
	position := (*component.PositionData)(entry.Component(component.Position))
	rect := (*component.RectData)(entry.Component(component.Rect))
	isBat := (*component.IsBatData)(entry.Component(component.IsBat))
	isBat.Value = false
	position.X = 0
	position.Y = 50
	rect.Height = 5
	rect.Width = float64(screenWidth)
}

func NewBottomBorder(w donburi.World) {
	screenWidth, screenHeight := ebiten.WindowSize()
	tb := w.Create(component.Rect, component.Position, component.IsBat)
	entry := w.Entry(tb)
	position := (*component.PositionData)(entry.Component(component.Position))
	rect := (*component.RectData)(entry.Component(component.Rect))
	isBat := (*component.IsBatData)(entry.Component(component.IsBat))
	isBat.Value = false
	position.X = 0
	position.Y = float64(screenHeight) - 50
	rect.Height = 5
	rect.Width = float64(screenWidth)
}
