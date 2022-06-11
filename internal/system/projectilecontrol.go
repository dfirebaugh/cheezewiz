package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/ecs/adapter"
	"math"
)

type ProjectileControl struct {
	ecs            adapter.Adapter
	projectilCache []archetype.Projectile
}

const projectileSpeed = 1.5

func NewProjectileContol(adapter adapter.Adapter) *ProjectileControl {
	return &ProjectileControl{
		ecs: adapter,
	}
}

func (p ProjectileControl) Update() {
	projectiles, ok := p.ecs.GetProjectiles()
	if !ok {
		return
	}
	for _, e := range projectiles {
		position := e.GetPosition()
		direction := e.GetDirection()

		vy := math.Sin(direction.Angle)
		vx := math.Cos(direction.Angle)

		vy *= projectileSpeed
		vx *= projectileSpeed

		position.Update(position.X-vx, position.Y-vy)
	}
}
