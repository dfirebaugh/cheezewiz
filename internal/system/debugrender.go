package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Debug struct {
	w world.World
}

func NewDebugRenderer(w world.World) Debug {
	return Debug{
		w: w,
	}
}

func (d Debug) Draw(screen *ebiten.Image) {
	if !config.Get().DebugEnabled {
		return
	}

	entities := d.w.Len()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f -- entities: %d\n", ebiten.CurrentFPS(), entities))

	if !config.Get().DebugCollidablesEnabled {
		return
	}
	query.Each(d.w, filter.GetAnimatables, func(h world.EntityHandle) {
		d.w.GetEntity(h).DebugDraw(screen)
	})
}
