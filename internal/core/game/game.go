// game is some stuff to make ebiten happy
//   (e.g. the Draw and  Update functions)
package game

import (
	"cheezewiz/internal/mediator"
	"image/color"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Width    int
	Height   int
	mediator mediator.Mediator
}

func New(screenWidth int, screenHeight int) *Game {
	game := &Game{
		Width:    screenWidth,
		Height:   screenHeight,
		mediator: mediator.New(),
	}
	return game
}

func (g *Game) Update() error {
	g.mediator.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	g.mediator.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / 2, outsideHeight / 2
}

func (g *Game) Run() {
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Cheezewiz")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
