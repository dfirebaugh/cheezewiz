package system

import (
	"cheezewiz/config"
	"cheezewiz/pkg/ecs"

	"github.com/sirupsen/logrus"
)

type WorldViewPortLocation struct {
	world ecs.World
}

func NewWorldViewPortLocation(w ecs.World) *WorldViewPortLocation {
	return &WorldViewPortLocation{
		world: w,
	}
}

func (w *WorldViewPortLocation) Update() {
	initialPlayer, err := ecs.FirstEntity[Player](w.world)
	if err != nil {
		logrus.Errorf("unable to find player: %s", err)
		return
	}
	playerPosition := initialPlayer.GetPosition()

	worldViewPort, err := ecs.FirstEntity[ViewPort](w.world)
	if err != nil {
		logrus.Errorf("viewport update: %s", err)
		return
	}
	worldViewPortPos := worldViewPort.GetPosition()

	worldViewPortPos.X = getWorldViewCenterLocation(playerPosition.X, config.Get().Window.Height) + 20
	worldViewPortPos.Y = 0
}

func getWorldViewCenterLocation(coordinate float64, windowDim int) float64 {
	return coordinate - float64(windowDim/config.Get().ScaleFactor)/2
}
