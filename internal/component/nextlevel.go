package component

import "github.com/yohamta/donburi"

type NextLevelData struct {
	CurrentLevel uint
}

var NextLevel = donburi.NewComponentType(InputDeviceData{})

func GetNextLevel(entry *donburi.Entry) *NextLevelData {
	return (*NextLevelData)(entry.Component(NextLevel))
}
