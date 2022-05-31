package component

import "github.com/yohamta/donburi"

type VelocityData struct {
	L, M float64
}

var Velocity = donburi.NewComponentType(VelocityData{})

func GetVelocity(entry *donburi.Entry) *VelocityData {
	return (*VelocityData)(entry.Component(Velocity))
}
