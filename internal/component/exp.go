package component

import "github.com/yohamta/donburi"

type ExpData struct {
	DesiredExp float64 `json:"desiredXp"`
	CurrentExp float64 `json:"currentXp"`
}

var Exp = donburi.NewComponentType(ExpData{})

func GetExp(e *donburi.Entry) *ExpData {
	return (*ExpData)(e.Component(Exp))
}
