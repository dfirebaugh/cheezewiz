package entity

import (
	"cheezewiz/internal/component"

	"github.com/yohamta/donburi"
)

var TimerTag = donburi.NewTag()

func MakeTimer(w donburi.World) {
	timer := w.Create(TimerTag, component.Position, component.Countdown)
	entry := w.Entry(timer)

	// position := (*component.PositionData)(entry.Component(component.Position))
	// position.Set(float64(constant.ScreenHeight/4)-20, 40)

	countdown := (*component.CountdownData)(entry.Component(component.Countdown))
	*countdown = component.CountdownData{
		CountDownInSec: 120,
	}
}
