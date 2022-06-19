package system

import (
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"
	"math"
)

type ProjectileControl struct {
	w world.World
}

const projectileSpeed = 1.5

func NewProjectileContol(w world.World) *ProjectileControl {
	return &ProjectileControl{
		w: w,
	}
}

func (p ProjectileControl) Update() {
	query.Each(p.w, filter.GetProjectiles, func(handle world.EntityHandle) {
		p.projectile(handle)
	})
}

func (p ProjectileControl) projectile(handle world.EntityHandle) {
	e := p.w.GetEntity(handle)

	position := e.GetPosition()
	direction := e.GetDirection()

	vy := math.Sin(direction.Angle)
	vx := math.Cos(direction.Angle)

	vy *= projectileSpeed
	vx *= projectileSpeed

	position.Update(position.X-vx, position.Y-vy)
}
