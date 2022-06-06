package system

import (
	"cheezewiz/examples/pong/internal/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Velocity struct {
	ballQuery *query.Query
}

func NewVelocity() *Velocity {
	return &Velocity{
		ballQuery: query.NewQuery(filter.Contains(
			component.Position,
			component.Radius,
			component.Velocity,
		))}
}

func (v Velocity) Update(w donburi.World) {
	v.ballQuery.EachEntity(w, func(entry *donburi.Entry) {
		velocity := component.GetVelocity(entry)
		position := component.GetPosition(entry)
		position.X += velocity.L
		position.Y += velocity.M
	})
}
