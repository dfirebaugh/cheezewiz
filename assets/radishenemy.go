package assets

import (
	"bytes"
	"cheezewiz/internal/constant"
	_ "embed"
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

//go:embed radishred.png
var radishEnemyRaw []byte

//go:embed radishblue.png
var radishEnemyBlueRaw []byte

//go:embed radishyellow.png
var radishEnemyYellowRaw []byte

var (
	redDecoded, _, _    = image.Decode(bytes.NewReader(radishEnemyRaw))
	blueDecoded, _, _   = image.Decode(bytes.NewReader(radishEnemyBlueRaw))
	yellowDecoded, _, _ = image.Decode(bytes.NewReader(radishEnemyYellowRaw))
	raddishGrid         = ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), redDecoded.Bounds().Dx(), redDecoded.Bounds().Dy())
	raddishSprites      = []*ganim8.Sprite{
		ganim8.NewSprite(ebiten.NewImageFromImage(redDecoded), raddishGrid.GetFrames("1-4", 1)),
		ganim8.NewSprite(ebiten.NewImageFromImage(blueDecoded), raddishGrid.GetFrames("1-4", 1)),
		ganim8.NewSprite(ebiten.NewImageFromImage(yellowDecoded), raddishGrid.GetFrames("1-4", 1)),
	}
)

func GetRaddishhWalking() (*ganim8.Sprite, *ganim8.Animation) {
	rnd := rand.Intn(len(raddishSprites))

	return raddishSprites[rnd], ganim8.NewAnimation(raddishSprites[rnd], 100*time.Millisecond, ganim8.Nop)
}
