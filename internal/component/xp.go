package component

import "github.com/yohamta/donburi"

type XPData struct {
	Value float64
}

var XP = donburi.NewComponentType(XPData{})

func GetXP(entry *donburi.Entry) *XPData {
	return (*XPData)(entry.Component(XP))
}
