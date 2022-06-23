package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/scenebuilder"
	"cheezewiz/internal/system"
	"cheezewiz/pkg/ebitenwrapper"
	"image/color"

	"github.com/sirupsen/logrus"
)

func initScene() {

}

func main() {
	logrus.SetLevel(logrus.ErrorLevel)

	c := config.Get()

	systems := []scenebuilder.System{
		mover{},
		collision{},
		clickSpawner{},
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
