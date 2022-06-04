package component

import (
	"github.com/yohamta/donburi"
)

type RigidBodyData struct {
	H float64
	W float64
	X float64
	Y float64
}

var RigidBody = donburi.NewComponentType(RigidBodyData{})

func GetRigidBody(entry *donburi.Entry) *RigidBodyData {
	return (*RigidBodyData)(entry.Component(RigidBody))
}
