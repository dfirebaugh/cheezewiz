package system

import (
	"cheezewiz/internal/world"
)

type Collision struct {
	w world.World
}

func NewCollision(w world.World) *Collision {
	return &Collision{
		w: w,
	}
}

func (c *Collision) Update() {
	// we don't need to evaluate collisions on every update
	// so let's slow it down a bit
	// if throttle.ShouldThrottle("collisionsystem", 5) {
	// 	return
	// }

	// collidables := []world.EntityHandle{}

	// c.w.EachEntity(func(handle world.EntityHandle) {
	// 	if !c.w.GetEntity(handle).HasTag(tag.Collidable) {
	// 		return
	// 	}

	// 	collidables = append(collidables, handle)
	// })
	// for _, handle := range collidables {
	// 	c.updateCollidable(handle, collidables)
	// }
}

func (c Collision) updateCollidable(id world.EntityHandle, collidableHandles []world.EntityHandle) {
	// collidable := c.w.GetEntity(id)
	// rb := collidable.GetRigidBody()
	// p := collidable.GetPosition()
	// state := collidable.GetState()
	// if state.GetCurrent() == component.DeathState {
	// 	return
	// }

	// ax := p.X - rb.L
	// ay := p.Y - rb.T
	// aw := rb.GetWidth()
	// ah := rb.GetHeight()

	// for _, handle := range collidableHandles {
	// 	// would be more efficient if we filter out things that this entity is collidable with
	// 	if !c.w.GetEntity(handle).HasTag(tag.Collidable) {
	// 		return
	// 	}

	// 	if id == handle {
	// 		return
	// 	}

	// 	collidableB := c.w.GetEntity(handle)
	// 	p2 := collidableB.GetPosition()
	// 	erb := collidableB.GetRigidBody()
	// 	bx := p2.X - erb.L
	// 	by := p2.Y - erb.T
	// 	bw := erb.GetWidth()
	// 	bh := erb.GetHeight()
	// 	if gamemath.Rect([]float64{ax, ay, aw, ah}).IsAxisAlignedCollision(gamemath.Rect([]float64{bx, by, bw, bh})) {
	// 		if rb.CollisionHandler == nil {
	// 			return
	// 		}
	// 		rb.CollisionHandler(c.w, collidableB)
	// 	}
	// }
}
