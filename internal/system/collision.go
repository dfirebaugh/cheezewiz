package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/throttle"
)

type Collision struct {
	world ecs.World
}

func NewCollision(world ecs.World) *Collision {
	return &Collision{
		world: world,
	}
}

func (c Collision) Update() {

	// we don't need to evaluate collisions on every update
	// so let's slow it down a bit
	if throttle.ShouldThrottle("collisionsystem", 5) {
		return
	}

	for id, Collidable := range ecs.FilterBy[archetype.Collidable](c.world) {
		c.updateCollidable(id, Collidable)
	}
}

func (c Collision) updateCollidable(id int, collidable archetype.Collidable) {
	rb := collidable.GetRigidBody()
	p := collidable.GetPosition()

	ax := p.X - rb.L
	ay := p.Y - rb.T
	aw := rb.GetWidth()
	ah := rb.GetHeight()

	for idB, collidableB := range ecs.FilterBy[archetype.Collidable](c.world) {
		if id == idB {
			continue
		}
		p2 := collidableB.GetPosition()
		erb := collidableB.GetRigidBody()
		bx := p2.X - erb.L
		by := p2.Y - erb.T
		bw := erb.GetWidth()
		bh := erb.GetHeight()
		if c.IsCollide([]float64{ax, ay, aw, ah}, []float64{bx, by, bw, bh}) {
			if rb.CollisionHandler == nil {
				return
			}
			rb.CollisionHandler(c.world, collidableB)
		}

	}
}

func (c Collision) IsCollide(a []float64, b []float64) bool {
	ax := a[0]
	ay := a[1]
	aw := a[2]
	ah := a[3]

	bx := b[0]
	by := b[1]
	bw := b[2]
	bh := b[3]

	return ax < bx+bw &&
		ax+aw > bx &&
		ay < by+bh &&
		ah+ay > by
}
