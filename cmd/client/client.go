package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/core/game"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.ErrorLevel)

	c := config.Get()
	game.New(c.Window.Height, c.Window.Width).Run()
}
