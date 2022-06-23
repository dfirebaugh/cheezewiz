package ebitenwrapper

import (
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type scene interface {
	Update()
	Draw(screen *ebiten.Image)
	Exit()
}

type Game struct {
	Width  int
	Height int
	Scene  scene
	ebiten.Game
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
}

func (g *Game) Update() error {
	g.Scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / g.WindowScale, outsideHeight / g.WindowScale
}

func (g *Game) Run() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc
		g.Exit()
	}()

	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle(g.WindowTitle)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Exit() {
	g.Exit()
}
