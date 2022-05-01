package entity

import (
	"cheezewiz/config"
	"cheezewiz/internal/ecs/component"

	"code.rocketnine.space/tslocum/brownboxbatman/asset"
	"code.rocketnine.space/tslocum/gohan"
)

func NewPlayer() gohan.Entity {
	player := gohan.NewEntity()

	c := config.Get()
	player.AddComponent(&component.Position{
		X: float64(c.Window.Width) / 4,
		Y: float64(c.Window.Height) / 4,
	})
	player.AddComponent(&component.Velocity{
		X: 0,
		Y: 0,
	})
	player.AddComponent(&component.Radius{
		Value: 15,
	})
	player.AddComponent(&component.Asset{Image: asset.ImgWhiteSquare})

	return player
}
