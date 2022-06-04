package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"
	"cheezewiz/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Collision struct {
	playerQuery *query.Query
	enemyQuery  *query.Query
}

func NewCollision() *Collision {
	return &Collision{
		playerQuery: query.NewQuery(filter.Contains(
			entity.PlayerTag,
		)),
		enemyQuery: query.NewQuery(filter.Contains(
			entity.EnemyTag,
		)),
	}
}

func (c *Collision) Update(w donburi.World) {
	c.playerQuery.EachEntity(w, func(entry *donburi.Entry) {
		playerPosition := component.GetPosition(entry)
		playerHealth := component.GetHealth(entry)
		c.enemyQuery.EachEntity(w, func(entry *donburi.Entry) {
			enemyPosition := component.GetPosition(entry)
			if c.IsCollide(component.RigidBodyData{
				H: float64(constant.SpriteSize) - 4,
				W: float64(constant.SpriteSize) - 4,
				X: playerPosition.X,
				Y: playerPosition.Y,
			}, component.RigidBodyData{
				H: float64(constant.SpriteSize) - 4,
				W: float64(constant.SpriteSize) - 4,
				X: enemyPosition.X,
				Y: enemyPosition.Y,
			}) {
				if playerHealth.HP > 0 {
					playerHealth.HP--
				}
			}
		})
	})
}

func (c *Collision) Draw(w donburi.World, screen *ebiten.Image) {
}

func (c *Collision) IsCollide(a component.RigidBodyData, b component.RigidBodyData) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
