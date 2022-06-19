package system

import (
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	w  world.World
	op *ebiten.DrawImageOptions
}

func NewRenderer(w world.World) Renderer {
	return Renderer{
		w:  w,
		op: &ebiten.DrawImageOptions{},
	}
}

func (r *Renderer) Update() {
	query.Each(r.w, filter.GetAnimatables, func(handle world.EntityHandle) {
		r.updateAnimatable(handle)
	})
}

func (r *Renderer) updateAnimatable(handle world.EntityHandle) {
	e := r.w.GetEntity(handle)
	e.IterFrame()
}

func (r Renderer) Draw(screen *ebiten.Image) {
	r.resetOcclusionMap()

	query.Each(r.w, filter.GetAnimatables, func(handle world.EntityHandle) {
		r.drawAnimatables(handle, screen)
	})
}

func (r *Renderer) drawAnimatables(handle world.EntityHandle, screen *ebiten.Image) {
	e := r.w.GetEntity(handle)
	r.op.GeoM.Reset()
	e.Draw(screen, r.op)
}

func (r *Renderer) resetOcclusionMap() {
	// r.occlusionMap = [][]archetype.Animatable{}

	// for y := 0; y < 1024; y++ {
	// 	row := []archetype.Animatable{}
	// 	for x := 0; x < 1024; x++ {
	// 		row = append(row, nil)
	// 	}
	// 	r.occlusionMap = append(r.occlusionMap, row)
	// }
}
