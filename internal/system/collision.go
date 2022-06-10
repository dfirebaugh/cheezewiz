package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/throttle"

	cache "github.com/Code-Hex/go-generics-cache"
)

type Collision struct {
	world           ecs.World
	collidableCache *cache.Cache[cacheKey, []archetype.Collidable]
	prevCount       int
}

func NewCollision(world ecs.World) *Collision {
	c := cache.New[cacheKey, []archetype.Collidable]()
	c.Set(collidable, ecs.FilterBy[archetype.Collidable](world))

	return &Collision{
		world:           world,
		collidableCache: c,
	}
}

func (c Collision) Update() {
	// we don't need to evaluate collisions on every update
	// so let's slow it down a bit
	if throttle.ShouldThrottle("collisionsystem", 5) {
		return
	}
	var collidables []archetype.Collidable
	var ok bool

	if c.prevCount != c.world.Count() {
		c.prevCount = c.world.Count()
		c.collidableCache.Set(collidable, ecs.FilterBy[archetype.Collidable](c.world))
	}

	if collidables, ok = c.collidableCache.Get(collidable); !ok {
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
