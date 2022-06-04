package system

import (
	"cheezewiz/examples/pong/internal/component"
	"cheezewiz/internal/input"

	"github.com/sedyh/mizu/pkg/engine"
)

type Player struct{}

func (p *Player) Update(w engine.World) {
	view := w.View(component.Pos{}, component.Rect{}, component.Vel{}, component.IsPlayer{})

	view.Each(func(entity engine.Entity) {
		var pos *component.Pos
		var vel *component.Vel
		var isPlayer *component.IsPlayer
		entity.Get(&pos, &vel, &isPlayer)

		if !isPlayer.Value {
			return
		}

		k := input.Keyboard{}

		if k.IsUpPressed() {
			pos.Y -= 4
		}

		if k.IsDownPressed() {
			pos.Y += 4
		}
	})
}
