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

type DebugRenderer struct{}

func (d DebugRenderer) Draw(screen *ebiten.Image) {
	if !config.Get().DebugEnabled {
		return
	}

	entities := world.Instance.Len()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f -- entities: %d\n", ebiten.CurrentFPS(), entities))

	if !config.Get().DebugCollidablesEnabled {
		return
	}
	query.Each(world.Instance, filter.GetAnimatables, func(h world.EntityHandle) {
		world.Instance.GetEntity(h).DebugDraw(screen)
	})
}
