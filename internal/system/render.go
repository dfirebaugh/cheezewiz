package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Renderable interface {
	Draw(screen *ebiten.Image)
}

type Render struct {
	count int
}

func NewRender() *Render {
	return &Render{}
}

func (r *Render) Update(w donburi.World) {
	r.count++
}

func (r Render) Draw(w donburi.World, screen *ebiten.Image) {
}
