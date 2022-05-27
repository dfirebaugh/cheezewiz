package system

import (
	"cheezewiz/examples/pong/internal/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Collision struct{}

func (c *Collision) Update(w engine.World) {
	ballView := w.View(component.Pos{}, component.Rad{})
	view := w.View(component.Pos{}, component.Rect{}, component.IsBat{})
	ballView.Each(func(entity engine.Entity) {
		var ballPos *component.Pos
		var vel *component.Vel
		var rad *component.Rad
		entity.Get(&ballPos, &vel, &rad)
		view.Each(func(entity2 engine.Entity) {
			if entity2.ID() == entity.ID() {
				return
			}
			var pos *component.Pos
			var rigidBody *component.Rect
			var isBat *component.IsBat
			entity2.Get(&pos, &rigidBody, &isBat)

			ballRB := component.RigidBody{
				X: ballPos.X,
				Y: ballPos.Y,
				W: rad.Value,
				H: rad.Value,
			}
			rb := component.RigidBody{
				X: pos.X,
				Y: pos.Y,
				H: rigidBody.Height,
				W: rigidBody.Width,
			}

			if c.IsCollide(ballRB, rb) {
				vel.M *= -1.01
				if isBat.Value {
					vel.L *= -1.01
				}
			}
		})
	})
}

func (c *Collision) IsCollide(a component.RigidBody, b component.RigidBody) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
