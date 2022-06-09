package collision

import (
	"cheezewiz/internal/tag"
	"cheezewiz/pkg/attackgroup"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type HandlerLabel string

const (
	PlayerCollisionLabel    HandlerLabel = "player"
	EnemyCollisionLabel     HandlerLabel = "enemy"
	BossCollisionLabel      HandlerLabel = "boss"
	RocketCollisionLabel    HandlerLabel = "rocket"
	JellyBeanCollisionLabel HandlerLabel = "jellybean"
)

var c = map[HandlerLabel]func(w donburi.World, e *donburi.Entry){
	EnemyCollisionLabel: func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(tag.Player) {
			// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
			attackgroup.AddPlayerDamage(e, 10, nil)
		}
		if e.Archetype().Layout().HasComponent(tag.Projectile) {
			logrus.Info("enemy collided with projectile")
			w.Remove(e.Entity())
		}
	},
	RocketCollisionLabel: func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(tag.Enemy) {
			logrus.Info("missile collided with enemy")
			// w.Remove(e.Entity())
			// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
			attackgroup.AddEnemyDamage(e, 10, nil)
		}
	},
	BossCollisionLabel: func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(tag.Player) {
			logrus.Info("collision with boss")
			// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
			attackgroup.AddPlayerDamage(e, 10, nil)
		}

	},
	PlayerCollisionLabel: func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(tag.Enemy) {
			logrus.Info("player collided with enemy")
			// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
			// attackgroup.AddEnemyDamage(e, 10, nil)
		}
		if e.Archetype().Layout().HasComponent(tag.JellyBean) {
			logrus.Info("player collided with enemy")
			w.Remove(e.Entity())
		}
	},
}

func GetCollisionHandler(label HandlerLabel) (func(w donburi.World, e *donburi.Entry), error) {
	if _, ok := c[label]; !ok {
		return nil, fmt.Errorf("could not find collision handler: %s", label)
	}

	return c[label], nil
}
