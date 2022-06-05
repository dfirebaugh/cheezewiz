package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/internal/component"
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type JellyBeanType uint

const (
	JBBlue JellyBeanType = iota
	JBGreen
	JBRainbow
)

var JellyBeanTag = donburi.NewTag()

func MakeJellyBean(w donburi.World, x float64, y float64) {
	b := w.Create(JellyBeanTag, component.XP, component.SpriteSheet, component.Position, component.RigidBody)
	entry := w.Entry(b)

	xp := (*component.XPData)(component.GetXP(entry))
	xp.Value = 1
	sprite := (*component.SpriteSheetData)(component.GetSpriteSheet(entry))
	position := (*component.PositionData)(component.GetPosition(entry))

	blue, _, _ := image.Decode(bytes.NewReader(assets.JellyBeanBlueRaw))
	green, _, _ := image.Decode(bytes.NewReader(assets.JellyBeanGreenRaw))
	rainbow, _, _ := image.Decode(bytes.NewReader(assets.JellyBeanRainbowRaw))

	switch (JellyBeanType)(rand.Intn(3)) {
	case JBBlue:
		sprite.IMG = ebiten.NewImageFromImage(blue)
	case JBRainbow:
		sprite.IMG = ebiten.NewImageFromImage(rainbow)
	default:
		sprite.IMG = ebiten.NewImageFromImage(green)
	}

	position.Set(x, y, float64(sprite.IMG.Bounds().Dx())/2, float64(sprite.IMG.Bounds().Dy())/2)

	rb := (*component.RigidBodyData)(component.GetRigidBody(entry))
	rb.SetBorder(float64(sprite.IMG.Bounds().Dx()), float64(sprite.IMG.Bounds().Dy()))
}
