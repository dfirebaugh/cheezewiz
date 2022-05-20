package scene

import (
	"cheezewiz/internal/ecs/component"
	"cheezewiz/internal/ecs/entity"
	"cheezewiz/internal/ecs/system"
	"cheezewiz/internal/input"
	"os"

	_ "embed"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Update()
	Render(screen *ebiten.Image)
	AttachLevel(*component.Level)
}

type Scene struct {
	Player  gohan.Entity
	Level   *component.Level
	Systems []System
}

//go:embed main.js
var main string

func Init() *Scene {
	s := &Scene{}
	s.Level = &component.Level{
		EntityMap: make(map[int]interface{}),
		Entities:  []int{},
	}

	s.Level.Add(entity.NewPlayer(input.Keyboard{}))
	s.AddSystem(&system.Animator{})
	s.AddSystem(&system.Movement{})
	s.AddSystem(&system.Control{})

	// v := vm.Build(s)
	// vm.Run(v, main)

	return s
}

func (s *Scene) AddSystem(sys System) {
	sys.AttachLevel(s.Level)
	s.Systems = append(s.Systems, sys)
}

func (s *Scene) Update() {
	for _, system := range s.Systems {
		system.Update()
	}

	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}
func (s *Scene) Draw(screen *ebiten.Image) {
	for _, system := range s.Systems {
		system.Render(screen)
	}
}

func (s *Scene) Exit() {
	os.Exit(0)
}
