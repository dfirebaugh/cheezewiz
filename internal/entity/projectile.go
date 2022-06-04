package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/yohamta/ganim8/v2"
)

var ProjectileTag = donburi.NewTag()

func MakeRocketProjectile(w donburi.World, x float64, y float64, dir float64) *donburi.Entry {
	b := w.Create(ProjectileTag, component.Position, component.SpriteSheet, component.Animation, component.Direction)

	entry := w.Entry(b)

	position := (*component.PositionData)(entry.Component(component.Position))

	animation := (*component.AnimationData)(entry.Component(component.Animation))

	direction := (*component.DirectionData)(entry.Component(component.Direction))

	position.X = x
	position.Y = y

	direction.Angle = dir

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.MissleRaw))

	grid := ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

	animation.Walk.Sprite = ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-4", 1))
	animation.Walk.Animation = ganim8.NewAnimation(animation.Walk.Sprite, 100*time.Millisecond, ganim8.Nop)

	return entry
}
