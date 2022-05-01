package input

import "github.com/hajimehoshi/ebiten/v2"

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)
}
func (Keyboard) IsDownPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)
}
func (Keyboard) IsLeftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
}
func (Keyboard) IsRightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)
}
