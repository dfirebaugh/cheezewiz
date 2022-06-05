package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"
	"cheezewiz/internal/input"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var PlayerTag = donburi.NewTag()

func MakePlayer(w donburi.World, controller input.PlayerInput) *donburi.Entry {
	b := w.Create(PlayerTag, component.Position, component.Animation, component.InputDevice, component.Health, component.Direction, component.ActorState, component.RigidBody)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	inputDevice := (*component.InputDeviceData)(entry.Component(component.InputDevice))
	inputDevice.Device = controller

	health := (*component.HealthAspect)(entry.Component(component.Health))

	health.HP = 100
	health.MAXHP = 100

	animation := (*component.AnimationData)(entry.Component(component.Animation))
	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.CheezeWizRaw))
	hurtIMGDecoded, _, _ := image.Decode(bytes.NewReader(assets.CheezeWizHurtRaw))

	grid := ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

	idleSprite := ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1", 1))
	walkingSprite := ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-3", 1))
	hurtSprite := ganim8.NewSprite(ebiten.NewImageFromImage(hurtIMGDecoded), grid.GetFrames("1-3", 1))

	direction := (*component.DirectionData)(entry.Component(component.Direction))
	direction.Angle = 0

	animation.Animations = map[component.ActorStateType]component.Action{
		component.IdleState: {
			Sprite:    idleSprite,
			Animation: ganim8.NewAnimation(idleSprite, 100*time.Millisecond, ganim8.Nop),
		},
		component.WalkingState: {
			Sprite:    walkingSprite,
			Animation: ganim8.NewAnimation(walkingSprite, 100*time.Millisecond, ganim8.Nop),
		},
		component.HurtState: {
			Sprite:    hurtSprite,
			Animation: ganim8.NewAnimation(hurtSprite, 100*time.Millisecond, ganim8.Nop),
		},
	}

	position.Set(
		float64(config.Get().Window.Height/config.Get().ScaleFactor)/2,
		float64(config.Get().Window.Height/config.Get().ScaleFactor)/2,
		float64(walkingSprite.Width())/2,
		float64(walkingSprite.Height())/2,
	)

	collisionHandler := func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(JellyBeanTag) {
			w.Remove(e.Entity())
		}
	}

	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))
	*rb = component.RigidBodyData{
		L:                config.Get().TileSize / 2,
		R:                config.Get().TileSize / 2,
		T:                config.Get().TileSize / 2,
		B:                config.Get().TileSize / 2,
		Name:             "projectile",
		CollisionHandler: collisionHandler,
	}

	return entry
}
