package scene

import (
	"cheezewiz/config"
	"cheezewiz/internal/attacks"
	"cheezewiz/internal/dentity"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/system"
	"cheezewiz/pkg/taskrunner"
	"os"
	"time"

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

	// System
	renderer := system.NewRender()
	collision := system.NewCollision()
	timer := system.NewTimer()
	exp := system.NewExpbar()
	aicontroller := system.NewEnemyControl()

	taskrunner.Add(time.Millisecond*800, attacks.CheeseMissile(world))
	addEntities(world)

	s := &Scene{
		world: world,
		systems: []System{
			renderer,
			system.NewPlayerControl(),
			timer,
			system.NewRegisterPlayer(),
			system.DamageBufferGroup{},
			aicontroller,
			collision,
			system.NewScheduler(loadWorld(level1).Events, world),
			system.NewWorldViewPortLocation(),
			system.NewProjectileContol(),
			exp,
		},
		drawables: []Drawable{
			collision,
			renderer,
			timer,
			// exp,
		},
	}

	return s
}

func addEntities(world donburi.World) {
	// entity.MakeExpBar(world)
	entity.MakeWorld(world)
	entity.MakeBackground(world)
	entity.MakeTimer(world)
	// entity.MakePlayer(world, input.Keyboard{})
	dentity.MakeRandEntity(
		world,
		[]string{
			"entities/jellybeangreen.entity.json",
			"entities/jellybeanpink.entity.json",
			"entities/jellybeanblue.entity.json",
			"entities/jellybeanrainbow.entity.json",
		},
		200,
		200,
	)
	dentity.MakeEntity(
		world,
		"entities/cheezewiz.entity.json",
		float64(config.Get().Window.Width/config.Get().ScaleFactor/2),
		float64(config.Get().Window.Height/config.Get().ScaleFactor/2),
	)
	// entity.MakeSlot(world)
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
