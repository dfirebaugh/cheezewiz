package component

import "github.com/yohamta/donburi"

type HealthAspect struct {
	MAXHP float64 `json:"max"`
	HP    float64 `json:"hp"`
}

var Health = donburi.NewComponentType(HealthAspect{})

func GetHealth(e *donburi.Entry) *HealthAspect {
	return (*HealthAspect)(e.Component(Health))
}
