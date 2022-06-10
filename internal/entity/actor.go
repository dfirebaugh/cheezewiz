package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	*component.Animation
	*component.State
	*component.Position
	*component.Health
	*component.RigidBody
}

func (e Actor) GetRigidBody() *component.RigidBody {
	return e.RigidBody
}
func (e Actor) GetFrame() *ebiten.Image {
	return e.Animation.Animation[e.State.GetCurrent()].GetFrame()
}
func (e Actor) GetPosition() *component.Position {
	return e.Position
}
func (e Actor) GetCurrentState() component.StateType {
	return e.State.GetCurrent()
}
func (e Actor) GetCurrent() *animation.Animation {
	return e.Animation.Animation[e.GetCurrentState()]
}
func (e Actor) IterFrame() {
	e.GetCurrent().IterFrame()
}
func (e Actor) GetHealth() *component.Health {
	return e.Health
}
func (e Actor) GetState() *component.State {
	return e.State
}
