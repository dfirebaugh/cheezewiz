package scene

import (
	"cheezewiz/internal/entity"
	"cheezewiz/internal/input"
	"cheezewiz/internal/system"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}
type Scene struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

func Init() *Scene {
	renderer := system.NewRender()
	collision := system.NewCollision()
	s := &Scene{
		world: donburi.NewWorld(),
		systems: []System{
			renderer,
			system.NewPlayerControl(),
		},
		drawables: []Drawable{
			collision,
			renderer,
		},
	}

	entity.MakePlayer(s.world, input.Keyboard{})

	return s
}

func (s *Scene) Update() {
	for _, sys := range s.systems {
		sys.Update(s.world)
	}
	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}
func (s *Scene) Draw(screen *ebiten.Image) {
	for _, sys := range s.drawables {
		sys.Draw(s.world, screen)
	}
}

func (s *Scene) Exit() {
	os.Exit(0)
}
