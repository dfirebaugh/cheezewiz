package component

import "github.com/yohamta/donburi"

type IsBatData struct {
	Value bool
}

var IsBat = donburi.NewComponentType(IsBatData{})

func GetIsBat(entry *donburi.Entry) *IsBatData {
	return (*IsBatData)(entry.Component(IsBat))
}
