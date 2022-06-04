package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/internal/component"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var SlotTag = donburi.NewTag()

// a slot contains UI elements for players
func MakeSlot(w donburi.World) {
	b := w.Create(SlotTag, component.SpriteSheet)
	entry := w.Entry(b)
	sprite := (*component.SpriteSheetData)(entry.Component(component.SpriteSheet))

	img, _, _ := image.Decode(bytes.NewReader(assets.CheezeWizSlotRaw))
	sprite.IMG = ebiten.NewImageFromImage(img)

}
