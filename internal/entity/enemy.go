package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var EnemyTag = donburi.NewTag()

func MakeEnemy(w donburi.World, x float64, y float64) *donburi.Entry {
	b := w.Create(EnemyTag, component.Position, component.Health, component.SpriteSheet, component.Animation, component.Direction, component.ActorState, component.RigidBody)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))

	health := (*component.HealthAspect)(entry.Component(component.Health))

	health.HP = 10

	position.Set(x, y, 15, 15)

	state := (*component.ActorStateData)(entry.Component(component.ActorState))
	state.Set(component.WalkingState)

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.RadishEnemyRaw))
	animation := (*component.AnimationData)(entry.Component(component.Animation))
	grid := ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())
	walkingSprite := ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-4", 1))

	animation.Animations = map[component.ActorStateType]component.Action{
		component.WalkingState: {
			Sprite:    walkingSprite,
			Animation: ganim8.NewAnimation(walkingSprite, 100*time.Millisecond, ganim8.Nop),
		},
	}

	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))
	rb.SetBorder(config.Get().TileSize, config.Get().TileSize)
	rb.Name = "radish"

	return entry
}

func MakeBossEnemy(w donburi.World, x float64, y float64, hp float64) *donburi.Entry {
	b := w.Create(EnemyTag, component.Position, component.Health, component.Animation, component.Direction, component.ActorState, component.RigidBody)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))

	health := (*component.HealthAspect)(entry.Component(component.Health))

	health.HP = 200

	position.Set(x, y, 15, 15)
	state := (*component.ActorStateData)(entry.Component(component.ActorState))
	state.Set(component.WalkingState)

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.CheezeBossRaw))

	grid := ganim8.NewGrid(int(constant.SpriteSize*2), int(constant.SpriteSize*2), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())
	animation := (*component.AnimationData)(entry.Component(component.Animation))
	walkingSprite := ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-5", 1))

	animation.Animations = map[component.ActorStateType]component.Action{
		component.WalkingState: {
			Sprite:    walkingSprite,
			Animation: ganim8.NewAnimation(walkingSprite, 100*time.Millisecond, ganim8.Nop),
		},
	}
	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))
	rb.Name = "boss"
	rb.SetBorder(config.Get().TileSize, config.Get().TileSize)

	// rb.CollisionHandler = func(e *donburi.Entry) {
	// 	if health.HP > 0 {
	// 		MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
	// 		// c.attack_handler.AddEnemyDamage(entry, 10, nil)
	// 	}
	// }

	return entry
}
