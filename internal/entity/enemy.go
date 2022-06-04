package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/internal/component"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var EnemyTag = donburi.NewTag()

func MakeEnemy(w donburi.World, x float64, y float64) *donburi.Entry {
	b := w.Create(EnemyTag, component.Position, component.Health, component.SpriteSheet, component.Animation)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))

	animation := (*component.AnimationData)(entry.Component(component.Animation))

	health := (*component.HealthAspect)(entry.Component(component.Health))

	health.HP = 10

	position.X = x
	position.Y = y

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.RadishEnemyRaw))

	grid := ganim8.NewGrid(32, 32, imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

	animation.Walk.Sprite = ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-3", 1))
	animation.Walk.Animation = ganim8.NewAnimation(animation.Walk.Sprite, 100*time.Millisecond, ganim8.Nop)

	return entry
}
