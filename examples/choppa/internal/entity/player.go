package entity

import (
	"bytes"
	"cheezewiz/examples/choppa/assets"
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/internal/input"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var PlayerTag = donburi.NewTag()

func NewPlayer(w donburi.World, controller input.PlayerInput) *donburi.Entry {
	b := w.Create(PlayerTag, component.Position, component.SpriteSheet, component.Direction, component.InputDevice)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	spriteSheet := (*component.SpriteSheetData)(entry.Component(component.SpriteSheet))
	facing := (*component.DirectionData)(entry.Component(component.Direction))

	inputDevice := (*component.InputDeviceData)(entry.Component(component.InputDevice))
	inputDevice.Device = controller

	facing.IsRight = true
	position.X = 0
	position.Y = 0

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.ChoppaRaw))
	spriteSheet.IMG = ebiten.NewImageFromImage(imgDecoded)

	return entry
}
