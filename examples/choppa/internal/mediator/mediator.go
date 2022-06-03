package mediator

import (
	"cheezewiz/examples/choppa/internal/entity"
	"cheezewiz/examples/choppa/internal/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Mediator struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

func New() *Mediator {
	renderer := system.NewRender()
	collision := system.NewCollision()
	m := &Mediator{
		world: donburi.NewWorld(),
		systems: []System{
			renderer,
			system.NewPlayer(),
			system.NewVelocity(),
			system.NewLifeSpan(),
			system.NewSpawner(),
			collision,
		},
		drawables: []Drawable{
			collision,
			renderer,
		},
	}
	entity.NewPlayer(m.world)
	return m
}

func (m Mediator) Update() {
	for _, s := range m.systems {
		s.Update(m.world)
	}
}

func (m Mediator) Draw(screen *ebiten.Image) {
	for _, s := range m.drawables {
		s.Draw(m.world, screen)
	}
}

func (Mediator) Exit() {}
