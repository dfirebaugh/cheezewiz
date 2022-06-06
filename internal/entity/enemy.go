package entity

import (
	"cheezewiz/assets"
	"cheezewiz/config"
	"cheezewiz/internal/component"

	"github.com/yohamta/donburi"
)

var EnemyTag = donburi.NewTag()

func MakeEnemy(w donburi.World, x float64, y float64) *donburi.Entry {
	b := w.Create(EnemyTag, component.Position, component.Health, component.SpriteSheet, component.Animation, component.Direction, component.ActorState, component.RigidBody)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	health := (*component.HealthAspect)(entry.Component(component.Health))
	state := (*component.ActorStateData)(entry.Component(component.ActorState))
	animation := (*component.AnimationData)(entry.Component(component.Animation))
	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))

	health.HP = 10
	position.Set(x, y, 15, 15)

	animation.Animations = map[component.ActorStateType]component.Action{
		component.WalkingState: component.ToAction(assets.GetRaddishhWalking()),
	}

	rb.SetBorder(config.Get().TileSize, config.Get().TileSize)
	state.Set(component.WalkingState)

	return entry
}

func MakeBossEnemy(w donburi.World, x float64, y float64, hp float64) *donburi.Entry {
	b := w.Create(EnemyTag, component.Position, component.Health, component.Animation, component.Direction, component.ActorState, component.RigidBody)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	animation := (*component.AnimationData)(entry.Component(component.Animation))
	health := (*component.HealthAspect)(entry.Component(component.Health))
	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))
	state := (*component.ActorStateData)(entry.Component(component.ActorState))

	health.HP = 200

	position.Set(x, y, 15, 15)
	state.Set(component.WalkingState)

	animation.Animations = map[component.ActorStateType]component.Action{
		component.WalkingState: component.ToAction(assets.GetCheeseBossWalking()),
	}
	rb.SetBorder(config.Get().TileSize, config.Get().TileSize)

	// rb.CollisionHandler = func(e *donburi.Entry) {
	// 	if health.HP > 0 {
	// 		MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
	// 		// c.attack_handler.AddEnemyDamage(entry, 10, nil)
	// 	}
	// }

	return entry
}
