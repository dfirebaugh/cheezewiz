package entity

import (
	"bytes"
	"cheezewiz/config"
	"cheezewiz/examples/sprite/internal/ecs/component"
	"cheezewiz/internal/input"
	"image"
	"log"
	"time"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/ganim8/v2/examples/assets/images"

	_ "image/png"
)

func NewPlayer() gohan.Entity {
	player := gohan.NewEntity()

	controller := input.Keyboard{}

	c := config.Get()
	player.AddComponent(&component.Position{
		X: float64(c.Window.Width) / 4,
		Y: float64(c.Window.Height) / 4,
	})
	player.AddComponent(&component.Velocity{
		X: 0,
		Y: 0,
	})

	grid := ganim8.NewGrid(100, 100, 500, 600)
	walkSprite := ganim8.NewSprite(
		ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_MONSTER_SLIME_BLUE)),
		grid.GetFrames("1-5", 5),
	)
	stillSprite := ganim8.NewSprite(
		ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_MONSTER_SLIME_BLUE)),
		grid.GetFrames("5", 5),
	)

	op := ganim8.DrawOpts(20, 20)
	op.SetOrigin(20, 20)

	player.AddComponent(&component.Animation{
		Walk:        ganim8.NewAnimation(walkSprite, 100*time.Millisecond, ganim8.Nop),
		Still:       ganim8.NewAnimation(stillSprite, 100*time.Millisecond, ganim8.Nop),
		Grid:        grid,
		WalkSprite:  walkSprite,
		StillSprite: stillSprite,
		SpriteSize:  c.TileSize,
		DrawOptions: op,
	})

	player.AddComponent(&component.Controller{
		Controller: controller,
	})

	player.AddComponent(&component.Movable{
		IsMoving: func() bool {
			return controller.IsLeftPressed() ||
				controller.IsRightPressed() ||
				controller.IsUpPressed() ||
				controller.IsDownPressed()
		},
	})

	return player
}

func bytes2Image(rawImage *[]byte) image.Image {
	img, format, error := image.Decode(bytes.NewReader(*rawImage))
	if error != nil {
		log.Fatal("Bytes2Image Failed: ", format, error)
	}
	return img
}
