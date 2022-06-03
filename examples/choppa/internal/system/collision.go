package system

import (
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/examples/choppa/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/exp/shiny/materialdesign/colornames"
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
				w.Remove(entry.Entity())
			}
		})
		c.fighterQuery.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)

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
				w.Remove(entry.Entity())
			}
		})
	})
}

func (c *Collision) Draw(w donburi.World, screen *ebiten.Image) {
	renderRB := false
	if !renderRB {
		return
	}
	c.playerProjectileQuery.EachEntity(w, func(entry *donburi.Entry) {
		pPosition := component.GetPosition(entry)

		c.chippaQuery.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)
			ebitenutil.DrawRect(screen, pPosition.X, pPosition.Y, 100, 5, colornames.Red100)
			ebitenutil.DrawRect(screen, position.X, position.Y, 32, 32, colornames.Red100)
		})
		c.fighterQuery.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)
			ebitenutil.DrawRect(screen, pPosition.X, pPosition.Y, 100, 5, colornames.Red100)
			ebitenutil.DrawRect(screen, position.X, position.Y, 32, 32, colornames.Red100)
		})
	})
}

func (c *Collision) IsCollide(a component.RigidBody, b component.RigidBody) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
