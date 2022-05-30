package system

import (
	"cheezewiz/examples/pongv2/internal/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Collision struct {
	query     *query.Query
	ballQuery *query.Query
}

func NewCollision() *Collision {
	return &Collision{
		query: query.NewQuery(filter.Contains(
			component.Position,
			component.Rect,
			component.IsBat,
		)),
		ballQuery: query.NewQuery(filter.Contains(
			component.Position,
			component.Radius,
			component.Velocity,
		))}
}

func (c *Collision) Update(w donburi.World) {
	c.ballQuery.EachEntity(w, func(entry *donburi.Entry) {
		ballPosition := component.GetPosition(entry)
		ballRadius := component.GetRadius(entry)
		velocity := component.GetVelocity(entry)

		c.query.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)
			rect := component.GetRect(entry)
			isBat := component.GetIsBat(entry)

			if c.IsCollide(component.RigidBody{
				X: ballPosition.X,
				Y: ballPosition.Y,
				W: ballRadius.Value,
				H: ballRadius.Value,
			}, component.RigidBody{
				X: position.X,
				Y: position.Y,
				W: rect.Width,
				H: rect.Height,
			}) {
				if isBat.Value {
					velocity.L *= -1.2
					return
				}
				velocity.M *= -1
			}
		})
	})
}

func (c *Collision) IsCollide(a component.RigidBody, b component.RigidBody) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
