package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/internal/ecs/adapter"
	"cheezewiz/pkg/gamemath"
	"cheezewiz/pkg/throttle"
)

type Collision struct {
	ecs adapter.Adapter
}

func NewCollision(adapter adapter.Adapter) *Collision {
	return &Collision{
		ecs: adapter,
	}
}

func (c *Collision) Update() {
	// we don't need to evaluate collisions on every update
	// so let's slow it down a bit
	if throttle.ShouldThrottle("collisionsystem", 5) {
		return
	}

	collidables, ok := c.ecs.GetCollidables()
	if !ok {
		return
	}

	for id, collidable := range collidables {
		c.updateCollidable(id, collidables, collidable)
	}
}

func (c Collision) updateCollidable(id int, collidables []archetype.Collidable, collidable archetype.Collidable) {
	rb := collidable.GetRigidBody()
	p := collidable.GetPosition()
	state := collidable.GetState()
	if state.GetCurrent() == component.DeathState {
		return
	}

	ax := p.X - rb.L
	ay := p.Y - rb.T
	aw := rb.GetWidth()
	ah := rb.GetHeight()

	for idB, collidableB := range collidables {
		if id == idB {
			continue
		}
		p2 := collidableB.GetPosition()
		erb := collidableB.GetRigidBody()
		bx := p2.X - erb.L
		by := p2.Y - erb.T
		bw := erb.GetWidth()
		bh := erb.GetHeight()
		if gamemath.Rect([]float64{ax, ay, aw, ah}).IsAxisAlignedCollision(gamemath.Rect([]float64{bx, by, bw, bh})) {
			if rb.CollisionHandler == nil {
				return
			}
			rb.CollisionHandler(c.ecs.GetWorld(), collidableB)
		}

	}
}
