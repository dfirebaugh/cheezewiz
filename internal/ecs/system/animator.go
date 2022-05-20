package system

import (
	"cheezewiz/internal/ecs/component"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// recurse through all entities
// if entity matches archetype for that system, apply the
//  system transforms to the entity
type Animatable interface {
	// GetAsset() *component.Asset
	GetAnimation() *component.Animation
	GetPosition() *component.Position
	GetMovable() *component.Movable
	GetDrawOptions() *ebiten.DrawImageOptions
}

type Animator struct {
	Level *component.Level
}

func (a *Animator) AttachLevel(lvl *component.Level) {
	a.Level = lvl
}
func (a Animator) Update() {
	now := time.Now()
	for _, id := range a.Level.Entities {
		if _, ok := a.Level.EntityMap[id].(Animatable); !ok {
			println("entity doens't match contract")
			continue
		}
		entity := a.Level.EntityMap[id].(Animatable)

		if entity.GetMovable().IsMoving() {
			entity.GetAnimation().Walk.Update(now.Sub(entity.GetAnimation().PrevUpdateTime))
		} else {
			entity.GetAnimation().Still.Update(now.Sub(entity.GetAnimation().PrevUpdateTime))
		}
		entity.GetAnimation().PrevUpdateTime = now
	}
}

func (a Animator) Render(screen *ebiten.Image) {
	for _, id := range a.Level.Entities {
		if _, ok := a.Level.EntityMap[id].(Animatable); !ok {
			println("entity doens't match contract")
			continue
		}

		entity := a.Level.EntityMap[id].(Animatable)

		entity.GetAnimation().DrawOptions.Reset()
		entity.GetAnimation().DrawOptions.SetPos(entity.GetPosition().X-entity.GetAnimation().SpriteSize, entity.GetPosition().Y-entity.GetAnimation().SpriteSize)

		if entity.GetMovable().IsMoving() {
			entity.GetAnimation().Walk.Draw(screen, entity.GetAnimation().DrawOptions)
		} else {
			entity.GetAnimation().Still.Draw(screen, entity.GetAnimation().DrawOptions)
		}
	}
}
