package system

import (
	"cheezewiz/examples/pong/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"
)

type Renderable interface {
	Draw(screen *ebiten.Image)
}

// When you need many sets of components
// in one system, you can use the views
type Render struct{}

// Render one frame
func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	// But choose the right entities yourself
	ballView := w.View(component.Pos{}, component.Rad{})
	ballView.Each(func(entity engine.Entity) {
		var pos *component.Pos
		var rad *component.Rad
		entity.Get(&pos, &rad)

		ebitenutil.DrawRect(
			screen,
			pos.X-rad.Value/2, pos.Y-rad.Value/2,
			rad.Value, rad.Value,
			colornames.Aliceblue,
		)
	})
	borderView := w.View(component.Pos{}, component.Rect{})
	borderView.Each(func(entity engine.Entity) {
		var pos *component.Pos
		var rect *component.Rect
		entity.Get(&pos, &rect)

		ebitenutil.DrawRect(
			screen,
			pos.X-rect.Width/2, pos.Y-rect.Height/2,
			rect.Width, rect.Height,
			colornames.Aliceblue,
		)
	})
}
