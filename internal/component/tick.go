package component

import "github.com/yohamta/donburi"

type TickData struct {
	Value uint
	EOL   uint
}

var Tick = donburi.NewComponentType(TickData{})

func GetTick(entry *donburi.Entry) *TickData {
	return (*TickData)(entry.Component(Tick))
}
