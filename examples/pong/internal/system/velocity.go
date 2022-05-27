package system

import (
	"cheezewiz/examples/pong/internal/component"

	"github.com/sedyh/mizu/pkg/engine"
)

// You can go through all entities that have a certain set of
// components specifying the requirements in the fields of the system
type Velocity struct {
	*component.Pos // Current entity position
	*component.Vel // Current entity velocity
}

// Apply velocity for each entity that has Pos and Vel
func (v *Velocity) Update(w engine.World) {
	// If they are registered components, they will not be nil
	v.Pos.X += v.Vel.L
	v.Pos.Y += v.Vel.M
}
