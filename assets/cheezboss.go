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

//go:embed cheezboss.png
var cheezeBossRaw []byte

var (
	imgDecoded, _, _           = image.Decode(bytes.NewReader(cheezeBossRaw))
	cheeseBossGrid             = ganim8.NewGrid(int(constant.SpriteSize*2), int(constant.SpriteSize*2), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())
	cheezeBossWalkingSprite    = ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), cheeseBossGrid.GetFrames("1-5", 1))
	cheezeBossWalkingAnimation = ganim8.NewAnimation(cheezeBossWalkingSprite, 100*time.Millisecond, ganim8.Nop)
)

func GetCheeseBossWalking() (*ganim8.Sprite, *ganim8.Animation) {
	return cheezeBossWalkingSprite, cheezeBossWalkingAnimation
}
