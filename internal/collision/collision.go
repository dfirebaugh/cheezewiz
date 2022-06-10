package collision

import (
	"cheezewiz/pkg/attackgroup"
	"cheezewiz/pkg/ecs"
	"fmt"

	"github.com/sirupsen/logrus"
)

type HandlerLabel string

type Projectile interface {
	GetProjectileTag() ecs.Tag
}

type Player interface {
	GetPlayerTag() ecs.Tag
}

type Enemy interface {
	GetEnemyTag() ecs.Tag
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
		if ecs.Is[Player](e) {
			logrus.Info("enemy collided with player")
			attackgroup.AddPlayerDamage(e, 10, nil)
		}
		if ecs.Is[Projectile](e) {
			logrus.Info("enemy collided with projectile")
			// remove projectile
			w.Remove(e)
		}
	},
	RocketCollisionLabel: func(w ecs.World, e any) {
		if ecs.Is[Enemy](e) {
			logrus.Info("rocket collided with enemy")
			attackgroup.AddEnemyDamage(e, 10, nil)
		}
	},
	BossCollisionLabel: func(w ecs.World, e any) {
	},
	PlayerCollisionLabel: func(w ecs.World, e any) {
		if ecs.Is[Enemy](e) {
			logrus.Info("player collided with enemy")
		}
	},
}

func GetCollisionHandler(label HandlerLabel) (func(w ecs.World, e any), error) {
	if _, ok := c[label]; !ok {
		return nil, fmt.Errorf("could not find collision handler: %s", label)
	}

	return c[label], nil
}
