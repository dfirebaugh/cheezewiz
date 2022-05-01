package system

import (
	"cheezewiz/internal/ecs/component"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type movement struct {
	Position *component.Position
	Velocity *component.Velocity
}

func NewMovement() gohan.System {
	return &movement{}
}

func (m *movement) Update(entity gohan.Entity) error {
	m.Position.X = m.Position.X + m.Velocity.X
	m.Position.Y = m.Position.Y + m.Velocity.Y

	return nil
}

func (m *movement) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
