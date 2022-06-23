package collision

import (
	"cheezewiz/internal/tag"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type World interface {
	Remove(handle entityHandle)
	GetEntity(handle entityHandle) entity
}
type entity interface {
	HasTag(tag.Tag) bool
}

type CollisionHandler func(w World, handle entityHandle)
type HandlerLabel string

type entityHandle uuid.UUID

const (
	PlayerCollisionLabel    HandlerLabel = "player"
	EnemyCollisionLabel     HandlerLabel = "enemy"
	BossCollisionLabel      HandlerLabel = "boss"
	RocketCollisionLabel    HandlerLabel = "rocket"
	JellyBeanCollisionLabel HandlerLabel = "jellybean"
)

var c = map[HandlerLabel]CollisionHandler{
	EnemyCollisionLabel: func(w World, h entityHandle) {
		e := w.GetEntity(h)
		if e.HasTag(tag.Player) {
			logrus.Info("enemy collided with player")
			// attackgroup.AddPlayerDamage(e, 10, nil)
		}
		if e.HasTag(tag.Projectile) {
			logrus.Info("enemy collided with projectile")
			// remove projectile
			w.Remove(h)
		}
	},
	RocketCollisionLabel: func(w World, h entityHandle) {
		// if ecs.Is[Enemy](e) {
		// 	logrus.Info("rocket collided with enemy")
		// 	attackgroup.AddEnemyDamage(e, 10, nil)
		// }
	},
	BossCollisionLabel: func(w World, h entityHandle) {
	},
	PlayerCollisionLabel: func(w World, h entityHandle) {
		// if ecs.Is[Enemy](e) {
		// 	logrus.Info("player collided with enemy")
		// }
	},
}

func GetCollisionHandler(label HandlerLabel) (func(w World, h entityHandle), error) {
	if _, ok := c[label]; !ok {
		return nil, fmt.Errorf("could not find collision handler: %s", label)
	}

	return c[label], nil
}
