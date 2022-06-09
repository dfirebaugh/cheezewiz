package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"
	"math"
)

type Projectile interface {
	GetDirection() *component.Direction
	GetPosition() *component.Position
	GetProjectileTag() ecs.Tag
}
type ProjectileControl struct {
	world ecs.World
}

const projectileSpeed = 1.5

func NewProjectileContol(w ecs.World) *ProjectileControl {
	return &ProjectileControl{
		world: w,
	}
}

func (p ProjectileControl) Update() {
	for _, e := range ecs.FilterBy[Projectile](p.world) {
		position := e.GetPosition()
		direction := e.GetDirection()

		vy := math.Sin(direction.Angle)
		vx := math.Cos(direction.Angle)

		vy *= projectileSpeed
		vx *= projectileSpeed

		position.Update(position.X-vx, position.Y-vy)
	}
}
