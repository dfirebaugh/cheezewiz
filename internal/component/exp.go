package component

import "github.com/yohamta/donburi"

type ExpData struct {
	DesiredExp float64
	CurrentExp float64
}

var Exp = donburi.NewComponentType(ExpData{})

func GetExp(e *donburi.Entry) *ExpData {
	return (*ExpData)(e.Component(Exp))
}
