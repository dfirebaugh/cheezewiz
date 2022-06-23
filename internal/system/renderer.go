package system

import (
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	op *ebiten.DrawImageOptions
}

func NewRenderer() Renderer {
	return Renderer{
		op: &ebiten.DrawImageOptions{},
	}
}

func (r *Renderer) Update() {
	query.Each(world.Instance, filter.GetAnimatables, func(handle world.EntityHandle) {
		r.updateAnimatable(handle)
	})
}

func (r *Renderer) updateAnimatable(handle world.EntityHandle) {
	e := world.Instance.GetEntity(handle)
	e.IterFrame()
}

func (r Renderer) Draw(screen *ebiten.Image) {
	query.Each(world.Instance, filter.GetAnimatables, func(handle world.EntityHandle) {
		r.drawAnimatables(handle, screen)
	})
}

func (r *Renderer) drawAnimatables(handle world.EntityHandle, screen *ebiten.Image) {
	e := world.Instance.GetEntity(handle)
	r.op.GeoM.Reset()
	e.Draw(screen, r.op)
}
