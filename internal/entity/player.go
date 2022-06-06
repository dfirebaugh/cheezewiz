package entity

import (
	"cheezewiz/assets"
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"

	"github.com/yohamta/donburi"
)

var PlayerTag = donburi.NewTag()

func MakePlayer(w donburi.World, controller input.PlayerInput) *donburi.Entry {
	b := w.Create(PlayerTag, component.Position, component.Animation, component.InputDevice, component.Health, component.Direction, component.ActorState, component.RigidBody)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	inputDevice := (*component.InputDeviceData)(entry.Component(component.InputDevice))
	inputDevice.Device = controller
	animation := (*component.AnimationData)(entry.Component(component.Animation))
	state := (*component.ActorStateData)(entry.Component(component.ActorState))
	direction := (*component.DirectionData)(entry.Component(component.Direction))
	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))

	health := (*component.HealthAspect)(entry.Component(component.Health))

	health.HP = 100
	health.MAXHP = 100

	direction.Angle = 0
	animation.Animations = map[component.ActorStateType]component.Action{
		component.IdleState:    component.ToAction(assets.GetCheezeWizIdle()),
		component.WalkingState: component.ToAction(assets.GetCheezeWizWalking()),
		component.HurtState:    component.ToAction(assets.GetCheezeWizHurt()),
	}

	position.Set(
		float64(config.Get().Window.Width/config.Get().ScaleFactor/2),
		float64(config.Get().Window.Height/config.Get().ScaleFactor/2),
		float64(config.Get().TileSize)/2,
		float64(config.Get().TileSize)/2,
	)
	collisionHandler := func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(component.JellyBeanTag) {
			w.Remove(e.Entity())
		}

		if e.Archetype().Layout().HasComponent(EnemyTag) {
			state.Set(component.HurtState)
		}
	}

	*rb = component.RigidBodyData{
		L:                config.Get().TileSize / 2,
		R:                config.Get().TileSize / 2,
		T:                config.Get().TileSize / 2,
		B:                config.Get().TileSize / 2,
		CollisionHandler: collisionHandler,
	}

	return entry
}
