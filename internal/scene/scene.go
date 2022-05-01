package scene

import (
	"cheezewiz/internal/ecs/entity"
	"cheezewiz/internal/ecs/system"
	"cheezewiz/internal/input"
	"os"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type scene struct {
	Player gohan.Entity
}

func Init() *scene {
	s := &scene{}

	s.addSystems()
	s.Player = entity.NewPlayer()

	return s
}

func (s *scene) Update() {
	err := gohan.Update()
	if err != nil {
		logrus.Error(err)
	}
	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}
func (s *scene) Draw(screen *ebiten.Image) {
	err := gohan.Draw(screen)
	if err != nil {
		panic(err)
	}
}
func (s *scene) addSystems() {
	gohan.AddSystem(system.NewPlayerControl(input.Keyboard{}))
	gohan.AddSystem(system.NewMovement())
	gohan.AddSystem(system.NewRenderer())
}
func (s *scene) Exit() {
	os.Exit(0)
}
