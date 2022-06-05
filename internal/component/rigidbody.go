package component

import (
	"github.com/yohamta/donburi"
)

type RigidBodyData struct {
	R                float64
	L                float64
	T                float64
	B                float64
	CollisionHandler func(w donburi.World, e *donburi.Entry)
	Name             string
}

var RigidBody = donburi.NewComponentType(RigidBodyData{})

func GetRigidBody(entry *donburi.Entry) *RigidBodyData {
	return (*RigidBodyData)(entry.Component(RigidBody))
}

func (r RigidBodyData) GetHeight() float64 {
	return r.R + r.L
}
func (r RigidBodyData) GetWidth() float64 {
	return r.T + r.B
}

func (r *RigidBodyData) SetBorder(wx float64, wy float64) {
	r.R = wx / 2
	r.L = wx / 2
	r.T = wy / 2
	r.B = wy / 2
}
