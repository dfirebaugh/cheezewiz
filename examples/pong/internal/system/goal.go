package system

import (
	"cheezewiz/examples/pong/internal/component"
	"cheezewiz/examples/pong/internal/entity"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

type Goal struct {
	Player   uint
	Computer uint
}

func (g *Goal) Update(w engine.World) {
	screenWidth, _ := ebiten.WindowSize()

	ballView := w.View(component.Pos{}, component.Rad{})
	ballView.Each(func(entity engine.Entity) {
		var pos *component.Pos
		entity.Get(&pos)

		if pos.X < 0 {
			g.Computer++
			w.RemoveEntity(entity)
			g.ResetBall(w)
		}
		if int(pos.X) > screenWidth {
			g.Player++
			w.RemoveEntity(entity)
			g.ResetBall(w)
		}
	})
}

func (g *Goal) Draw(w engine.World, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d:%d", g.Player, g.Computer))
}

func (g *Goal) ResetBall(w engine.World) {
	screenWidth, screenHeight := ebiten.WindowSize()
	w.AddEntities(&entity.Ball{
		Pos: component.Pos{
			X: float64(screenWidth) / 2,
			Y: float64(screenHeight) / 2,
		},
		Vel: component.Vel{
			L: 2,
			M: 2,
		},
		Rad: component.Rad{
			Value: 5,
		}},
	)
}
