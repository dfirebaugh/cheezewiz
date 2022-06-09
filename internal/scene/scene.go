package scene

import (
	"cheezewiz/internal/entity"
	"cheezewiz/internal/system"
	"cheezewiz/pkg/ecs"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	// Update(w donburi.World)
	Update()
}

type Drawable interface {
	// Draw(w donburi.World, screen *ebiten.Image)
	Draw(screen *ebiten.Image)
}
type Scene struct {
	// world     donburi.World
	world     ecs.World
	systems   []System
	drawables []Drawable
}

// const level1 string = "level1.json"

func Init(level string) *Scene {
	// World
	// world := donburi.NewWorld()
	w := ecs.NewWorld()

	// System
	// renderer := system.NewRender()
	// collision := system.NewCollision()
	// timer := system.NewTimer()
	// exp := system.NewExpbar()
	// aicontroller := system.NewEnemyControl()

	// taskrunner.Add(time.Millisecond*800, attacks.CheeseMissile(world))
	// addEntities(world)
	// entity.MakePlayer(w)
	entity.MakeEntity(w, "entities/cheezewiz.entity.json", 0, 0)
	// entity.MakeEntity(w, "entities/radishred.entity.json", 0, 0)
	renderer := system.NewRenderer(w)
	s := &Scene{
		world: w,
		systems: []System{
			renderer,
			system.Controller{World: w},
			// 	system.NewPlayerControl(),
			// 	timer,
			// 	system.NewRegisterPlayer(),
			// 	system.DamageBufferGroup{},
			// 	aicontroller,
			// 	collision,
			system.NewScheduler(loadWorld(level).Events, w),
			// 	system.NewWorldViewPortLocation(),
			// 	system.NewProjectileContol(),
			// 	exp,
		},
		drawables: []Drawable{
			// 	collision,
			renderer,
			// 	timer,
			// 	// exp,
		},
	}

	return s
}

// func addEntities(world donburi.World) {
// 	// entity.MakeExpBar(world)
// 	entity.MakeWorld(world)
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
// }

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
