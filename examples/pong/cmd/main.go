package main

import (
	"cheezewiz/examples/pong/internal/component"
	"cheezewiz/examples/pong/internal/entity"
	"cheezewiz/examples/pong/internal/system"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

// Main scene, you can use that for settings or main menu
func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Pos{},
		component.Vel{},
		component.Rad{},
		component.Rect{},
		component.IsPlayer{},
		component.IsBat{},
	)
	g.AddEntities(w)
	w.AddSystems(
		&system.Velocity{},
		&system.Render{},
		&system.Collision{},
		&system.Goal{
			Player:   0,
			Computer: 0,
		},
		&system.AI{},
		&system.Player{},
	)
}

func (g *Game) AddEntities(w engine.World) {
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
		&entity.Border{
			Rect: component.Rect{
				Height: 5,
				Width:  float64(screenWidth) * 2,
			},
			Pos: component.Pos{
				X: 0,
				Y: float64(screenHeight - 75),
			},
			IsBat: component.IsBat{
				Value: false,
			},
		},
		&entity.Border{
			Rect: component.Rect{
				Height: 5,
				Width:  float64(screenWidth) * 2,
			},
			Pos: component.Pos{
				X: 0,
				Y: float64(75),
			},
			IsBat: component.IsBat{
				Value: false,
			},
		},

		&entity.Bat{
			Pos: component.Pos{
				X: 20,
				Y: float64(screenHeight) / 2,
			},
			Vel: component.Vel{
				L: 0,
				M: 0,
			},
			Rect: component.Rect{
				Height: 50,
				Width:  10,
			},
			IsPlayer: component.IsPlayer{
				Value: true,
			},
			IsBat: component.IsBat{
				Value: true,
			},
		},
		&entity.Bat{
			Pos: component.Pos{
				X: float64(screenWidth) - 20,
				Y: float64(screenHeight) / 2,
			},
			Vel: component.Vel{
				L: 0,
				M: 0,
			},
			Rect: component.Rect{
				Height: 50,
				Width:  10,
			},
			IsPlayer: component.IsPlayer{
				Value: false,
			},
			IsBat: component.IsBat{
				Value: true,
			},
		},
	)
}

func main() {
	g := engine.NewGame(&Game{})
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
