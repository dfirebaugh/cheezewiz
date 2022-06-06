package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"
	"cheezewiz/pkg/taskrunner"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var ProjectileTag = donburi.NewTag()

type attackGroup interface {
	AddEnemyDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry)
}

func MakeRocketProjectile(w donburi.World, x float64, y float64, dir float64, attacks attackGroup) *donburi.Entry {
	b := w.Create(ProjectileTag, component.Position, component.Animation, component.Direction, component.ActorState, component.RigidBody, component.Tick)

	entry := w.Entry(b)

	position := (*component.PositionData)(entry.Component(component.Position))
	tick := (*component.TickData)(entry.Component(component.Tick))
	tick.Creation = time.Now()

	animation := (*component.AnimationData)(entry.Component(component.Animation))

	direction := (*component.DirectionData)(entry.Component(component.Direction))

	direction.Angle = dir

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.MissleRaw))

	grid := ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

	walkingSprite := ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-4", 1))

	state := (*component.ActorStateData)(entry.Component(component.ActorState))

	state.Set(component.WalkingState)
	animation.Animations = map[component.ActorStateType]component.Action{
		component.WalkingState: {
			Sprite:    walkingSprite,
			Animation: ganim8.NewAnimation(walkingSprite, 100*time.Millisecond, ganim8.Nop),
		},
	}

	position.Set(x, y,
		float64(walkingSprite.W())/2,
		float64(walkingSprite.H())/2,
	)

	rb := (*component.RigidBodyData)(entry.Component(component.RigidBody))
	rb.SetBorder(config.Get().TileSize/2, config.Get().TileSize/2)
	rb.CollisionHandler = func(w donburi.World, e *donburi.Entry) {
		if e.Archetype().Layout().HasComponent(EnemyTag) {
			// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
			w.Remove(entry.Entity())
			attacks.AddEnemyDamage(e, 10, nil)
		}
	}

	tick.Interval = time.Second * 5

	tick.EOL = time.Now().Add(time.Second * 5)
	tick.EOLEvent = func() {
		w.Remove(entry.Entity())
	}

	taskrunner.Add(tick.Interval, func() {
		if time.Since(tick.EOL) > 0 {
			tick.EOLEvent()
		}
	})

	return entry
}
