package entity

import (
	"bytes"
	"cheezewiz/config"
	"cheezewiz/internal/ecs/component"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/ganim8/v2/examples/assets/images"

	_ "image/png"
)

type Player struct {
	Position   *component.Position
	Velocity   *component.Velocity
	Animation  *component.Animation
	Movable    *component.Movable
	Controller *component.Controller

	Op *ebiten.DrawImageOptions
}

func (p Player) GetPosition() *component.Position {
	return p.Position
}
func (p Player) GetVelocity() *component.Velocity {
	return p.Velocity
}
func (p Player) GetAnimation() *component.Animation {
	return p.Animation
}
func (p Player) GetMovable() *component.Movable {
	return p.Movable
}
func (p Player) GetDrawOptions() *ebiten.DrawImageOptions {
	return p.Op
}
func (p Player) GetController() *component.Controller {
	return p.Controller
}

func NewPlayer(controller component.PlayerInput) *Player {
	c := config.Get()

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

	return &Player{
		Position: &component.Position{
			X: float64(c.Window.Width) / 4,
			Y: float64(c.Window.Height) / 4,
		},
		Velocity: &component.Velocity{
			X: 0,
			Y: 0,
		},
		Animation: &component.Animation{
			Walk:        ganim8.NewAnimation(walkSprite, 100*time.Millisecond, ganim8.Nop),
			Still:       ganim8.NewAnimation(stillSprite, 100*time.Millisecond, ganim8.Nop),
			Grid:        grid,
			WalkSprite:  walkSprite,
			StillSprite: stillSprite,
			SpriteSize:  c.SpriteSize,
			DrawOptions: op,
		},
		Movable: &component.Movable{
			IsMoving: func() bool {
				return controller.IsLeftPressed() ||
					controller.IsRightPressed() ||
					controller.IsUpPressed() ||
					controller.IsDownPressed()
			},
		},
		Op: &ebiten.DrawImageOptions{},
		Controller: &component.Controller{
			Controller: controller,
		},
	}
}

func bytes2Image(rawImage *[]byte) image.Image {
	img, format, error := image.Decode(bytes.NewReader(*rawImage))
	if error != nil {
		log.Fatal("Bytes2Image Failed: ", format, error)
	}
	return img
}
