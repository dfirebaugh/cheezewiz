package system

import (
	"cheezewiz/examples/choppa/internal/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Velocity struct {
	query *query.Query
}

func NewVelocity() *Velocity {
	return &Velocity{
		query: query.NewQuery(filter.Contains(
			component.Position,
			component.Velocity,
		)),
	}
}

func (v Velocity) Update(w donburi.World) {
	v.query.EachEntity(w, func(entry *donburi.Entry) {
		velocity := component.GetVelocity(entry)
		position := component.GetPosition(entry)
		position.X += velocity.L
		position.Y += velocity.M
	})
}
