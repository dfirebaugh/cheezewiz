package scene

import (
	"cheezewiz/internal/vm"
	"os"

	_ "embed"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Scene struct {
	Player gohan.Entity
}

//go:embed main.js
var mainRaw string

func Init() *Scene {
	s := &Scene{}
	v := vm.Build(s)
	vm.Run(v, string(mainRaw))

	return s
}

func (s *Scene) SetPlayer(player gohan.Entity) {
	s.Player = player
}

func (s *Scene) Update() {
	err := gohan.Update()
	if err != nil {
		logrus.Error(err)
	}
	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}
func (s *Scene) Draw(screen *ebiten.Image) {
	err := gohan.Draw(screen)
	if err != nil {
		panic(err)
	}
}

func (s *Scene) Exit() {
	os.Exit(0)
}
