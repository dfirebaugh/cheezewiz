package component

type RigidBody struct {
	R float64 `json:"r"`
	L float64 `json:"l"`
	T float64 `json:"t"`
	B float64 `json:"b"`
	// CollisionHandler      func(w world.World, e uuid.UUID)
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

func (r *RigidBody) SetCollisionHandler(label interface{}) {
	// var l collision.HandlerLabel
	// var ok bool

	// logrus.Info(string(label.(collision.HandlerLabel)))
	// if l, ok = label.(collision.HandlerLabel); !ok || l == "" {
	// 	logrus.Errorf("not a defined collision handler %s", label)
	// 	r.CollisionHandler = nil
	// 	return
	// }

	// handler, err := collision.GetCollisionHandler(l)
	// if err != nil {
	// 	logrus.Warnf("not able to set collision handler %s -- %s", label, err)
	// 	r.CollisionHandler = nil
	// 	return
	// }

	// r.CollisionHandler = handler
}
