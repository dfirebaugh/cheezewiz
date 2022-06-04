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

const level1 string = "level1.json"

func Init() *Scene {
	// World
	world := donburi.NewWorld()

	loadWorld(level1)

	// System
	renderer := system.NewRender()
	collision := system.NewCollision()
	timer := system.NewTimer()
	damageGroup := system.NewDamagebufferGroup()
	aicontroller := system.NewEnemyControl()

	s := &Scene{
		world: world,
		systems: []System{
			renderer,
			system.NewPlayerControl(),
			timer,
			system.NewRegisterPlayer(),
			damageGroup,
			aicontroller,
			collision,
		},
		drawables: []Drawable{
			collision,
			renderer,
			timer,
		},
	}

	addEntities(world)

	return s
}

func addEntities(world donburi.World) {
	entity.MakeBackground(world)
	entity.MakeTimer(world)
	entity.MakePlayer(world, input.Keyboard{})
	entity.MakeEnemy(world, 50, 50)
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
