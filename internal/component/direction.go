package component

import "github.com/yohamta/donburi"

type DirectionData struct {
	IsRight bool    `json:"isRight"`
	Angle   float64 `json:"angle"`
}

var Direction = donburi.NewComponentType(DirectionData{})

func GetDirection(entry *donburi.Entry) *DirectionData {
	return (*DirectionData)(entry.Component(Direction))
}
