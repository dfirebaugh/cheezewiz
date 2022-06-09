package collision

import (
	"cheezewiz/pkg/ecs"
	"fmt"
)

type HandlerLabel string

type ProjectileTag struct{}
type Projectile interface {
	GetProjectileTag() ProjectileTag
}

type PlayerTag struct{}
type Player interface {
	GetPlayerTag() PlayerTag
}

type EnemyTag struct{}
type Enemy interface {
	GetEnemyTag() EnemyTag
}

const (
	PlayerCollisionLabel    HandlerLabel = "player"
	EnemyCollisionLabel     HandlerLabel = "enemy"
	BossCollisionLabel      HandlerLabel = "boss"
	RocketCollisionLabel    HandlerLabel = "rocket"
	JellyBeanCollisionLabel HandlerLabel = "jellybean"
)

var c = map[HandlerLabel]func(w ecs.World, e any){
	EnemyCollisionLabel: func(w ecs.World, e any) {
		if ecs.IsType[Player](e) {
			println("enemy collided with player")
		}
		// if e.Archetype().Layout().HasComponent(tag.Player) {
		// 	// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
		// 	attackgroup.AddPlayerDamage(e, 10, nil)
		// }
		// if e.Archetype().Layout().HasComponent(tag.Projectile) {
		// 	logrus.Info("enemy collided with projectile")
		// 	// w.Remove(e.Entity())
		// }
	},
	RocketCollisionLabel: func(w ecs.World, e any) {
		if ecs.IsType[Enemy](w) {
			println("rocket collided with enemy")
		}
		// if e.Archetype().Layout().HasComponent(tag.Enemy) {
		// 	logrus.Info("missile collided with enemy")
		// 	// w.Remove(e.Entity())
		// 	// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
		// 	attackgroup.AddEnemyDamage(e, 10, nil)
		// }
	},
	BossCollisionLabel: func(w ecs.World, e any) {
		// if e.Archetype().Layout().HasComponent(tag.Player) {
		// 	logrus.Info("collision with boss")
		// 	// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
		// 	attackgroup.AddPlayerDamage(e, 10, nil)
		// }

	},
	PlayerCollisionLabel: func(w ecs.World, e any) {
		if ecs.IsType[Enemy](w) {
			println("player collided with enemy")
		}
		// if e.Archetype().Layout().HasComponent(tag.Enemy) {
		// 	logrus.Info("player collided with enemy")
		// 	// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
		// 	// attackgroup.AddEnemyDamage(e, 10, nil)
		// }
		// if e.Archetype().Layout().HasComponent(tag.JellyBean) {
		// 	logrus.Info("player collided with enemy")
		// 	// w.Remove(e.Entity())
		// }
	},
}

func GetCollisionHandler(label HandlerLabel) (func(w ecs.World, e any), error) {
	if _, ok := c[label]; !ok {
		return nil, fmt.Errorf("could not find collision handler: %s", label)
	}

	return c[label], nil
}
