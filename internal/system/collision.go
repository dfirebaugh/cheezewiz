package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"
	"cheezewiz/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type attackMediator interface {
	AddPlayerDamage(destination *donburi.Entry, amount float64, origin *donburi.Entry)
	AddEnemyDamage(destination *donburi.Entry, amount float64, origin *donburi.Entry)
}

type Collision struct {
	playerQuery     *query.Query
	enemyQuery      *query.Query
	projectileQuery *query.Query
	attack_handler  attackMediator
	jellyBeanQuery  *query.Query
}

func NewCollision(attack_handler attackMediator) *Collision {
	return &Collision{
		playerQuery: query.NewQuery(filter.Contains(
			entity.PlayerTag,
		)),
		enemyQuery: query.NewQuery(filter.Contains(
			entity.EnemyTag,
		)),
		projectileQuery: query.NewQuery(filter.Contains(
			entity.ProjectileTag,
		)),
		jellyBeanQuery: query.NewQuery(filter.Contains(entity.JellyBeanTag)),
		attack_handler: attack_handler,
	}
}

func (c *Collision) Update(w donburi.World) {
	c.projectileQuery.EachEntity(w, func(prjentry *donburi.Entry) {
		projectilePos := component.GetPosition(prjentry)
		c.enemyQuery.EachEntity(w, func(entry *donburi.Entry) {
			enemyPosition := component.GetPosition(entry)
			enemyHealth := component.GetHealth(entry)
			if c.IsCollide(component.RigidBodyData{
				H: float64(constant.SpriteSize) - 4,
				W: float64(constant.SpriteSize) - 4,
				X: projectilePos.X,
				Y: projectilePos.Y,
			}, component.RigidBodyData{
				H: float64(constant.SpriteSize) - 4,
				W: float64(constant.SpriteSize) - 4,
				X: enemyPosition.X,
				Y: enemyPosition.Y,
			}) {
				if enemyHealth.HP > 0 {
					c.attack_handler.AddEnemyDamage(entry, 10, nil)
				}
			}
		})
	})

	c.playerQuery.EachEntity(w, func(pentry *donburi.Entry) {
		playerPosition := component.GetPosition(pentry)
		playerHealth := component.GetHealth(pentry)
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
					c.attack_handler.AddPlayerDamage(pentry, 1, nil)
				}
			}
		})

		c.jellyBeanQuery.EachEntity(w, func(entry *donburi.Entry) {
			position := component.GetPosition(entry)
			xp := component.GetXP(entry)
			if c.IsCollide(component.RigidBodyData{
				H: float64(constant.SpriteSize) - 4,
				W: float64(constant.SpriteSize) - 4,
				X: playerPosition.X,
				Y: playerPosition.Y,
			}, component.RigidBodyData{
				H: float64(constant.SpriteSize) - 4,
				W: float64(constant.SpriteSize) - 4,
				X: position.X,
				Y: position.Y,
			}) {
				// trigger sopmething about XP
				logrus.Infof("player gains %d xp", xp.Value)
				w.Remove(entry.Entity())
			}
		})
	})
}

func (c *Collision) Draw(w donburi.World, screen *ebiten.Image) {
}

func (c *Collision) IsCollide(a component.RigidBodyData, b component.RigidBodyData) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
