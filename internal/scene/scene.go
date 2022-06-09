package scene

import (
	"cheezewiz/config"
	"cheezewiz/internal/attacks"
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/system"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/taskrunner"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Update()
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}
type Scene struct {
	world     ecs.World
	systems   []System
	drawables []Drawable
}

// const level1 string = "level1.json"

func Init(level string) *Scene {
	// World
	w := ecs.NewWorld()
	taskrunner.Add(time.Millisecond*800, attacks.CheeseMissile(w))
	addEntities(w)

	// System
	renderer := system.NewRenderer(w)

	s := &Scene{
		world: w,
		systems: []System{
			renderer,
			system.NewCollision(w),
			system.Controller{World: w},
			system.NewEnemyControl(w),
			system.NewScheduler(loadWorld(level).Events, w),
			system.NewWorldViewPortLocation(w),
			system.DamageBufferGroup{World: w},
			system.NewProjectileContol(w),
		},
		drawables: []Drawable{
			renderer,
		},
	}

	return s
}

func addEntities(world ecs.World) {
	// 	// entity.MakeExpBar(world)
	world.Add(&entity.ViewPort{
		Position: &component.Position{},
	})
	entity.MakeEntity(world, "entities/cheezewiz.entity.json",
		float64(config.Get().Window.Width/config.Get().ScaleFactor/2),
		float64(config.Get().Window.Height/config.Get().ScaleFactor/2))

	// 	entity.MakeBackground(world)
	// 	entity.MakeTimer(world)
	// 	// entity.MakePlayer(world, input.Keyboard{})
	// 	dentity.MakeRandEntity(
	// 		world,
	// 		[]string{
	// 			"entities/jellybeangreen.entity.json",
	// 			"entities/jellybeanpink.entity.json",
	// 			"entities/jellybeanblue.entity.json",
	// 			"entities/jellybeanrainbow.entity.json",
	// 		},
	// 		200,
	// 		200,
	// 	)
	// 	dentity.MakeEntity(
	// 		world,
	// 		"entities/cheezewiz.entity.json",
	// 		float64(config.Get().Window.Width/config.Get().ScaleFactor/2),
	// 		float64(config.Get().Window.Height/config.Get().ScaleFactor/2),
	// 	)
	// 	// entity.MakeSlot(world)
}

func (s *Scene) Update() {
	for _, sys := range s.systems {
		sys.Update()
	}
	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}
func (s *Scene) Draw(screen *ebiten.Image) {
	for _, sys := range s.drawables {
		sys.Draw(screen)
	}
}

func (s *Scene) Exit() {
	os.Exit(0)
}
