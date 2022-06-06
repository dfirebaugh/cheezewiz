package assets

import (
	"bytes"
	"cheezewiz/internal/constant"
	_ "embed"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

//go:embed cheezewiz.png
var cheezeWizRaw []byte

//go:embed cheezewiz-damaged.png
var cheezeWizHurtRaw []byte

//go:embed cheezewiz.slot.png
var CheezeWizSlotRaw []byte

var (
	cwWalking, _, _ = image.Decode(bytes.NewReader(cheezeWizRaw))
	cwHurt, _, _    = image.Decode(bytes.NewReader(cheezeWizHurtRaw))

	cheezeWizGrid = ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

	idleSprite    = ganim8.NewSprite(ebiten.NewImageFromImage(cwWalking), cheezeWizGrid.GetFrames("1", 1))
	walkingSprite = ganim8.NewSprite(ebiten.NewImageFromImage(cwWalking), cheezeWizGrid.GetFrames("1-3", 1))
	hurtSprite    = ganim8.NewSprite(ebiten.NewImageFromImage(cwHurt), cheezeWizGrid.GetFrames("1-3", 1))
)

func GetCheezeWizWalking() (*ganim8.Sprite, *ganim8.Animation) {
	return idleSprite, ganim8.NewAnimation(walkingSprite, 100*time.Millisecond, ganim8.Nop)
}
func GetCheezeWizIdle() (*ganim8.Sprite, *ganim8.Animation) {
	return walkingSprite, ganim8.NewAnimation(idleSprite, 100*time.Millisecond, ganim8.Nop)
}
func GetCheezeWizHurt() (*ganim8.Sprite, *ganim8.Animation) {
	return hurtSprite, ganim8.NewAnimation(hurtSprite, 100*time.Millisecond, ganim8.Nop)
}
