package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/pkg/ecs"
	"math"
)

type ProjectileControl struct {
	world          ecs.World
	projectilCache []archetype.Projectile
}

const projectileSpeed = 1.5

func NewProjectileContol(w ecs.World) *ProjectileControl {
	return &ProjectileControl{
		world: w,
	}
}

func (p ProjectileControl) Update() {
	if p.projectilCache == nil {
		p.projectilCache = ecs.FilterBy[archetype.Projectile](p.world)
	}
	for _, e := range p.projectilCache {
		position := e.GetPosition()
		direction := e.GetDirection()

		vy := math.Sin(direction.Angle)
		vx := math.Cos(direction.Angle)

		vy *= projectileSpeed
		vx *= projectileSpeed

		position.Update(position.X-vx, position.Y-vy)
	}
}
