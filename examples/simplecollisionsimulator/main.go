package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/scenebuilder"
	"cheezewiz/internal/system"
	"cheezewiz/internal/tag"
	"cheezewiz/internal/world"
	"cheezewiz/pkg/ebitenwrapper"
	"image/color"

	"github.com/sirupsen/logrus"
)

func initScene() {
	entity.MakeWithTags(world.Instance, "entities/cheezewiz.entity.json",
		float64(config.Get().Window.Width/config.Get().ScaleFactor/2),
		float64(config.Get().Window.Height/config.Get().ScaleFactor/2), []tag.Tag{tag.Player, tag.Animatable, tag.Collidable})
}

func main() {
	logrus.SetLevel(logrus.ErrorLevel)

	c := config.Get()

	systems := []scenebuilder.System{
		mover{},
		collision{},
		clickSpawner{},
		system.PlayerController{},
	}

	drawables := []scenebuilder.Drawable{
		system.DebugRenderer{},
		hud{},
		renderer{},
	}

	game := &ebitenwrapper.Game{
		Scene:           scenebuilder.New(systems, drawables, initScene),
		WindowTitle:     c.Title,
		WindowScale:     c.ScaleFactor,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: color.NRGBA{0x00, 0x40, 0x80, 0xff},
	}

	game.Run()
}
