package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type ProjectileControl struct {
	query *query.Query
}

const projectileSpeed = 1.5

func NewProjectileContol() *ProjectileControl {
	return &ProjectileControl{
		query: query.NewQuery(filter.Contains(entity.ProjectileTag)),
	}
}

func (p ProjectileControl) Update(w donburi.World) {
	p.query.EachEntity(w, func(e *donburi.Entry) {
		position := component.GetPosition(e)
		direction := component.GetDirection(e)

		vy := math.Sin(direction.Angle)
		vx := math.Cos(direction.Angle)

		vy *= projectileSpeed
		vx *= projectileSpeed

		position.Update(position.X-vx, position.Y-vy)
	})
}
