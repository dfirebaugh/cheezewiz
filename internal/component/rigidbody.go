package component

import (
	"cheezewiz/internal/collision"

	"github.com/sirupsen/logrus"
)

type RigidBody struct {
	R                     float64 `json:"r"`
	L                     float64 `json:"l"`
	T                     float64 `json:"t"`
	B                     float64 `json:"b"`
	CollisionHandler      collision.CollisionHandler
	CollisionHandlerLabel string `json:"collisionLabel"`
}

func (r RigidBody) GetHeight() float64 {
	return r.R + r.L
}
func (r RigidBody) GetWidth() float64 {
	return r.T + r.B
}

func (r *RigidBody) SetBorder(wx float64, wy float64) {
	r.R = wx / 2
	r.L = wx / 2
	r.T = wy / 2
	r.B = wy / 2
}

// func (r RigidBody) ToRect() gamemath.Rect {
// 	return gamemath.Rect([]float64{r.})
// }

func (r *RigidBody) SetCollisionHandler(label interface{}) {
	l, ok := label.(collision.HandlerLabel)
	if !ok || l == "" {
		logrus.Errorf("not a defined collision handler %s", label)
		r.CollisionHandler = nil
		return
	}

	handler, err := collision.GetCollisionHandler(l)
	if err != nil {
		logrus.Warnf("not able to set collision handler %s -- %s", label, err)
		r.CollisionHandler = nil
		return
	}

	r.CollisionHandler = handler
}
