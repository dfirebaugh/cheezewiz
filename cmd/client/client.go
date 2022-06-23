package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/scene"
	"cheezewiz/internal/scenebuilder"
	"cheezewiz/internal/system"
	"cheezewiz/pkg/ebitenwrapper"
	"image/color"

	"github.com/sirupsen/logrus"
)

func initScene() {}

func main() {
	logrus.SetLevel(logrus.ErrorLevel)
	c := config.Get()

	renderer := system.NewRenderer()
	systems := []scenebuilder.System{
		system.PlayerController{},
		&system.EnemyController{},
		system.NewScheduler(scene.LoadStressTest().Events),
		&renderer,
	}

	drawables := []scenebuilder.Drawable{
		system.DebugRenderer{},
		renderer,
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
