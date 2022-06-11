package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/ecs/adapter"

	"github.com/sirupsen/logrus"
)

type WorldViewPortLocation struct {
	ecs adapter.Adapter
}

func NewWorldViewPortLocation(adapter adapter.Adapter) *WorldViewPortLocation {
	return &WorldViewPortLocation{
		ecs: adapter,
	}
}

func (w *WorldViewPortLocation) Update() {
	initialPlayer, err := w.ecs.FirstPlayer()
	if err != nil {
		logrus.Errorf("unable to find player: %s", err)
		return
	}

	viewPort, err := w.ecs.FirstViewPort()
	if err != nil {
		logrus.Errorf("viewport update: %s", err)
		return
	}
	playerPosition := initialPlayer.GetPosition()
	worldViewPortPos := viewPort.GetPosition()

	worldViewPortPos.X = getWorldViewCenterLocation(playerPosition.X, config.Get().Window.Height) + 20
	worldViewPortPos.Y = 0
}

func getWorldViewCenterLocation(coordinate float64, windowDim int) float64 {
	return coordinate - float64(windowDim/config.Get().ScaleFactor)/2
}
