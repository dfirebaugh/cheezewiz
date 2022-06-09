package collision

import (
	"cheezewiz/pkg/attackgroup"
	"cheezewiz/pkg/ecs"
	"fmt"
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
		if ecs.IsType[Player](e) {
			println("enemy collided with player")
			attackgroup.AddPlayerDamage(e, 10, nil)
		}
	},
	RocketCollisionLabel: func(w ecs.World, e any) {
		if ecs.IsType[Enemy](w) {
			println("rocket collided with enemy")
			attackgroup.AddEnemyDamage(e, 10, nil)
		}
	},
	BossCollisionLabel: func(w ecs.World, e any) {
	},
	PlayerCollisionLabel: func(w ecs.World, e any) {
		if ecs.IsType[Enemy](w) {
			println("player collided with enemy")
		}
	},
}

func GetCollisionHandler(label HandlerLabel) (func(w ecs.World, e any), error) {
	if _, ok := c[label]; !ok {
		return nil, fmt.Errorf("could not find collision handler: %s", label)
	}

	return c[label], nil
}
