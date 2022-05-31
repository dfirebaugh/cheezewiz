package component

import "github.com/yohamta/donburi"

type DirectionData struct {
	IsRight bool
}

var Direction = donburi.NewComponentType(DirectionData{})

func GetDirection(entry *donburi.Entry) *DirectionData {
	return (*DirectionData)(entry.Component(Direction))
}
