package entity

import (
	"cheezewiz/examples/pongv2/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var BallTag = donburi.NewTag()

func NewBall(w donburi.World) {
	screenWidth, screenHeight := ebiten.WindowSize()

	b := w.Create(BallTag, component.Position, component.Velocity, component.Radius)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	velocity := (*component.VelocityData)(entry.Component(component.Velocity))
	radius := (*component.RadiusData)(entry.Component(component.Radius))
	position.X = float64(screenWidth / 2)
	position.Y = float64(screenHeight / 2)
	radius.Value = 5
	velocity.L = 2
	velocity.M = 2
}
