package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/constant"

	"github.com/yohamta/donburi"
)

var TimerTag = donburi.NewTag()

func MakeTimer(w donburi.World) {
	timer := w.Create(TimerTag, component.Position, component.Countdown)
	entry := w.Entry(timer)

	position := (*component.PositionData)(entry.Component(component.Position))

	*position = component.PositionData{
		X: float64(constant.ScreenHeight/2) - 20,
		Y: 20,
	}

	countdown := (*component.CountdownData)(entry.Component(component.Countdown))

	*countdown = component.CountdownData{
		CountDownInSec: 120,
	}
}
