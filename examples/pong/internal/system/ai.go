package system

import (
	"cheezewiz/examples/pong/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type AI struct{}

func (a *AI) Update(w engine.World) {
	_, screenHeight := ebiten.WindowSize()

	ballView := w.View(component.Pos{}, component.Rad{})
	view := w.View(component.Pos{}, component.Rect{}, component.Vel{}, component.IsPlayer{})

	ballView.Each(func(entity engine.Entity) {
		var ballPos *component.Pos
		entity.Get(&ballPos)

		view.Each(func(entity2 engine.Entity) {
			var pos *component.Pos
			var vel *component.Vel
			var isPlayer *component.IsPlayer
			entity2.Get(&pos, &vel, &isPlayer)

			if isPlayer.Value {
				return
			}

			if int(ballPos.Y) < int(pos.Y) {
				if pos.Y <= 0 {
					return
				}
				pos.Y += -2
			}
			if int(ballPos.Y) > int(pos.Y) {
				if int(pos.Y) >= screenHeight {
					return
				}
				pos.Y += 2
			}

		})
	})
}
