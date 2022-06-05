package component

import (
	"time"

	"github.com/yohamta/donburi"
)

type TickData struct {
	Interval  time.Duration
	Creation  time.Time
	EOL       time.Time
	TickEvent func()
	EOLEvent  func()
}

var Tick = donburi.NewComponentType(TickData{})

func GetTick(entry *donburi.Entry) *TickData {
	return (*TickData)(entry.Component(Tick))
}
