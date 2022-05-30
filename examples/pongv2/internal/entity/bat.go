package entity

import (
	"cheezewiz/examples/pongv2/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var EnemyTag = donburi.NewTag()
var PlayerTag = donburi.NewTag()

func NewPlayer(w donburi.World) {
	_, screenHeight := ebiten.WindowSize()

	b := w.Create(PlayerTag, component.Position, component.Velocity, component.Rect, component.IsBat)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	velocity := (*component.VelocityData)(entry.Component(component.Velocity))
	rect := (*component.RectData)(entry.Component(component.Rect))
	isBat := (*component.IsBatData)(entry.Component(component.IsBat))
	position.X = 15
	position.Y = float64(screenHeight / 2)
	rect.Height = 60
	rect.Width = 10
	velocity.L = 2
	velocity.M = 2
	isBat.Value = true

}

func NewEnemy(w donburi.World) {
	screenWidth, screenHeight := ebiten.WindowSize()

	b := w.Create(EnemyTag, component.Position, component.Velocity, component.Rect, component.IsBat)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	velocity := (*component.VelocityData)(entry.Component(component.Velocity))
	rect := (*component.RectData)(entry.Component(component.Rect))
	isBat := (*component.IsBatData)(entry.Component(component.IsBat))
	isBat.Value = true
	position.X = float64(screenWidth) - 30
	position.Y = float64(screenHeight / 2)
	rect.Height = 60
	rect.Width = 10
	velocity.L = 2
	velocity.M = 2
}
