package component

import (
	"cheezewiz/internal/collision"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type RigidBodyData struct {
	R                     float64 `json:"r"`
	L                     float64 `json:"l"`
	T                     float64 `json:"t"`
	B                     float64 `json:"b"`
	CollisionHandler      func(w donburi.World, e *donburi.Entry)
	CollisionHandlerLabel string `json:"collisionLabel"`
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

func (r *RigidBodyData) SetCollisionHandler(label interface{}) {
	var l collision.HandlerLabel
	var ok bool

	logrus.Info(string(label.(collision.HandlerLabel)))
	if l, ok = label.(collision.HandlerLabel); !ok || l == "" {
		logrus.Error("not a defined collision handler")
		return
	}

	handler, err := collision.GetCollisionHandler(l)
	if err != nil {
		logrus.Errorf("not able to set collision handler %s -- %s", label, err)
		r.CollisionHandler = func(w donburi.World, e *donburi.Entry) {}
		return
	}

	r.CollisionHandler = handler
}
