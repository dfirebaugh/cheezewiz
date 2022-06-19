package collision

// type HandlerLabel string

// type entity interface {
// 	HasTag(struct{}) bool
// }

// // type World interface {
// // 	Remove(handle ecs.EntityHandle)
// // 	GetEntity(handle ecs.EntityHandle)
// // }

// const (
// 	PlayerCollisionLabel    HandlerLabel = "player"
// 	EnemyCollisionLabel     HandlerLabel = "enemy"
// 	BossCollisionLabel      HandlerLabel = "boss"
// 	RocketCollisionLabel    HandlerLabel = "rocket"
// 	JellyBeanCollisionLabel HandlerLabel = "jellybean"
// )

// var c = map[HandlerLabel]func(w ecs.World, h ecs.EntityHandle){
// 	EnemyCollisionLabel: func(w ecs.World, h ecs.EntityHandle) {
// 		e := w.GetEntity(h)
// 		if e.HasTag(tag.Player) {
// 			logrus.Info("enemy collided with player")
// 			attackgroup.AddPlayerDamage(e, 10, nil)
// 		}
// 		if e.HasTag(tag.Projectile) {
// 			logrus.Info("enemy collided with projectile")
// 			// remove projectile
// 			w.Remove(e)
// 		}
// 	},
// 	RocketCollisionLabel: func(w ecs.World, h ecs.EntityHandle) {
// 		// if ecs.Is[Enemy](e) {
// 		// 	logrus.Info("rocket collided with enemy")
// 		// 	attackgroup.AddEnemyDamage(e, 10, nil)
// 		// }
// 	},
// 	BossCollisionLabel: func(w ecs.World, h ecs.EntityHandle) {
// 	},
// 	PlayerCollisionLabel: func(w ecs.World, h ecs.EntityHandle) {
// 		// if ecs.Is[Enemy](e) {
// 		// 	logrus.Info("player collided with enemy")
// 		// }
// 	},
// }

// func GetCollisionHandler(label HandlerLabel) (func(w ecs.World, h ecs.EntityHandle), error) {
// 	if _, ok := c[label]; !ok {
// 		return nil, fmt.Errorf("could not find collision handler: %s", label)
// 	}

// 	return c[label], nil
// }
