package input

import "github.com/hajimehoshi/ebiten/v2"

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyW)
}
func (Keyboard) IsDownPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyS)
}
func (Keyboard) IsLeftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyA)
}
func (Keyboard) IsRightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyD)
}
