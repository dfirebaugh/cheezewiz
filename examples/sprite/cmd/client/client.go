package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/mediator"
	"cheezewiz/pkg/ebitenwrapper"
	"image/color"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	c := config.Get()
	game := &ebitenwrapper.Game{
		Mediator:        mediator.New(),
		WindowTitle:     c.Title,
		WindowScale:     2,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: color.NRGBA{0x00, 0x40, 0x80, 0xff},
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc
		game.Exit()
	}()

	game.Run()
}
