package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/core/game"
)

func main() {
	c := config.Get()
	game.New(c.Window.Height, c.Window.Width).Run()
}
