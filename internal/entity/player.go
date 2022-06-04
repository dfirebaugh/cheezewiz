package entity

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var PlayerTag = donburi.NewTag()

func MakePlayer(w donburi.World, controller input.PlayerInput) *donburi.Entry {
	b := w.Create(PlayerTag, component.Position, component.SpriteSheet, component.Animation, component.InputDevice, component.Health, component.Direction, component.State)
	entry := w.Entry(b)
	position := (*component.PositionData)(entry.Component(component.Position))
	// spriteSheet := (*component.SpriteSheetData)(entry.Component(component.SpriteSheet))
	inputDevice := (*component.InputDeviceData)(entry.Component(component.InputDevice))
	inputDevice.Device = controller

	animation := (*component.AnimationData)(entry.Component(component.Animation))

	health := (*component.HealthAspect)(entry.Component(component.Health))

	health.HP = 80
	health.MAXHP = 100

	position.X = float64(200)
	position.Y = float64(200)
	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.CheezeWizRaw))

	grid := ganim8.NewGrid(32, 32, imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

	animation.Idle.Sprite = ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1", 1))
	animation.Idle.Animation = ganim8.NewAnimation(animation.Idle.Sprite, 100*time.Millisecond, ganim8.Nop)
	animation.Walk.Sprite = ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1-3", 1))
	animation.Walk.Animation = ganim8.NewAnimation(animation.Walk.Sprite, 100*time.Millisecond, ganim8.Nop)

	return entry
}
