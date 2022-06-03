package system

import (
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/examples/choppa/internal/entity"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Collision struct {
	playerProjectileQuery *query.Query
	fighterQuery          *query.Query
	chippaQuery           *query.Query
}

func NewCollision() *Collision {
	return &Collision{
		playerProjectileQuery: query.NewQuery(filter.Contains(
			entity.ProjectileTag,
		)),
		fighterQuery: query.NewQuery(filter.Contains(
			entity.FighterTag,
		)),
		chippaQuery: query.NewQuery(filter.Contains(
			entity.ChippaTag,
		))}
}

func (c *Collision) Update(w donburi.World) {
	c.playerProjectileQuery.EachEntity(w, func(entry *donburi.Entry) {
		pPosition := component.GetPosition(entry)

		c.chippaQuery.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)
			alive := component.GetAlive(entry)

			if c.IsCollide(component.RigidBody{
				X: pPosition.X,
				Y: pPosition.Y,
				W: 100,
				H: 5,
			}, component.RigidBody{
				X: position.X,
				Y: position.Y,
				W: 32,
				H: 32,
			}) {
				alive.IsAlive = false
			}
		})
		c.fighterQuery.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)
			alive := component.GetAlive(entry)

			if c.IsCollide(component.RigidBody{
				X: pPosition.X,
				Y: pPosition.Y,
				W: 100,
				H: 5,
			}, component.RigidBody{
				X: position.X,
				Y: position.Y,
				W: 32,
				H: 32,
			}) {
				alive.IsAlive = false
			}
		})
	})
}

func (c *Collision) IsCollide(a component.RigidBody, b component.RigidBody) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
