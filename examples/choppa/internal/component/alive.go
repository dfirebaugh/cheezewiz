package component

import "github.com/yohamta/donburi"

type AliveData struct {
	IsAlive bool
}

var IsAlive = donburi.NewComponentType(AliveData{})

func GetAlive(entry *donburi.Entry) *AliveData {
	return (*AliveData)(entry.Component(IsAlive))
}
