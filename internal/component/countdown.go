package component

import (
	"github.com/yohamta/donburi"
)

type CountdownData struct {
	CountDownInSec uint32 `json:"countDown"`
}

var Countdown = donburi.NewComponentType(CountdownData{})

func GetCountdown(entry *donburi.Entry) *CountdownData {
	return (*CountdownData)(entry.Component(Countdown))
}
