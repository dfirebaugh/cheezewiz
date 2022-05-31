package entity

import (
	"cheezewiz/examples/choppa/internal/component"

	"github.com/yohamta/donburi"
)

var ProjectileTag = donburi.NewTag()

func MakeProjectile(w donburi.World, origin *component.PositionData) {
	p := w.Create(ProjectileTag, component.Position, component.Velocity, component.Tick)

	entry := w.Entry(p)
	position := (*component.PositionData)(entry.Component(component.Position))
	velocity := (*component.VelocityData)(entry.Component(component.Velocity))
	tick := (*component.TickData)(entry.Component(component.Tick))

	*velocity = component.VelocityData{
		L: 5,
		M: 0,
	}

	tick.Value = 0
	tick.EOL = 50

	position.X = origin.X
	position.Y = origin.Y
}
