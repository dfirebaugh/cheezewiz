package mediator

import (
	"cheezewiz/examples/sprite/internal/scene"
	"fmt"
	"os"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type iscene interface {
	Draw(screen *ebiten.Image)
	Update()
}

type Mediator struct {
	scene iscene
}

func New() Mediator {
	return Mediator{
		scene: scene.Init(),
	}
}

func (m Mediator) Update() {
	m.scene.Update()
}

func (m Mediator) Draw(dst *ebiten.Image) {
	ebitenutil.DebugPrint(dst, fmt.Sprintf("# of entities: %d\n#renders per loop:%d", gohan.CurrentEntities(), gohan.CurrentDraws()))

	m.scene.Draw(dst)
}

func (m Mediator) Exit() {
	// do something on exit?
	os.Exit(0)
}
