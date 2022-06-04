package component

import "github.com/yohamta/donburi"

type HealthAspect struct {
	HP float64
}

var Health = donburi.NewComponentType(HealthAspect{HP: 100})

func GetHealth(e *donburi.Entry) *HealthAspect {
	return (*HealthAspect)(e.Component(Health))
}
