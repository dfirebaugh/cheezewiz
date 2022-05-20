package system

import (
	"cheezewiz/internal/ecs/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type Movable interface {
	GetPosition() *component.Position
	GetVelocity() *component.Velocity
}

type Movement struct {
	Level *component.Level
}

func (m *Movement) AttachLevel(lvl *component.Level) {
	m.Level = lvl
}
func (m Movement) Update() {
	for _, id := range m.Level.Entities {
		if _, ok := m.Level.EntityMap[id].(Movable); !ok {
			continue
		}
		entity := m.Level.EntityMap[id].(Movable)

		entity.GetPosition().X = entity.GetPosition().X + entity.GetVelocity().X
		entity.GetPosition().Y = entity.GetPosition().Y + entity.GetVelocity().Y
	}
}
func (m Movement) Render(screen *ebiten.Image) {}
