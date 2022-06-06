package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GamePad struct {
	ID ebiten.GamepadID
}

func (gp GamePad) IsUpPressed() bool {
	return ebiten.IsGamepadButtonPressed(gp.ID, ebiten.GamepadButton10)
}

func (gp GamePad) IsDownPressed() bool {
	return ebiten.IsGamepadButtonPressed(gp.ID, ebiten.GamepadButton12)
}

func (gp GamePad) IsLeftJustPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(gp.ID, ebiten.GamepadButton13)
}
func (gp GamePad) IsLeftPressed() bool {
	return ebiten.IsGamepadButtonPressed(gp.ID, ebiten.GamepadButton13)
}
func (gp GamePad) IsRightJustPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(gp.ID, ebiten.GamepadButton11)
}
func (gp GamePad) IsRightPressed() bool {
	return ebiten.IsGamepadButtonPressed(gp.ID, ebiten.GamepadButton11)
}

func (gp GamePad) IsPrimaryAtkPressed() bool {
	return ebiten.IsGamepadButtonPressed(gp.ID, ebiten.GamepadButton0)
}

func (gp GamePad) IsPrimaryAtkJustPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(gp.ID, ebiten.GamepadButton0)
}
