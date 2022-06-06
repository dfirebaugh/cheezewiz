package main

import (
	"cheezewiz/examples/pong/internal/entity"
	"cheezewiz/examples/pong/internal/system"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

func (g *Game) Setup() {
	g.world = donburi.NewWorld()
	score := system.NewScore()
	g.systems = []System{
		system.NewVelocity(),
		system.NewCollision(),
		system.NewPlayer(),
		system.NewAI(),
		score,
	}
	g.drawables = []Drawable{
		system.NewRender(),
		score,
	}

	g.AddEntities()

}

func (g *Game) AddEntities() {
	entity.NewBall(g.world)
	entity.NewTopBorder(g.world)
	entity.NewBottomBorder(g.world)
	entity.NewPlayer(g.world)
	entity.NewEnemy(g.world)
}

func (g *Game) Update() error {
	for _, s := range g.systems {
		s.Update(g.world)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Clear()
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := &Game{}
	g.Setup()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
