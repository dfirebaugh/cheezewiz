package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	*component.Animation
	*component.ActorState
	*component.Position
	*component.Health
	*component.RigidBody
}

func (e Actor) GetRigidBody() *component.RigidBody {
	return e.RigidBody
}
func (e Actor) GetFrame() *ebiten.Image {
	return e.Animation.Animation[e.ActorState.GetCurrent()].GetFrame()
}
func (e Actor) GetPosition() *component.Position {
	return e.Position
}
func (e Actor) GetState() component.ActorStateType {
	return e.ActorState.GetCurrent()
}
func (e Actor) GetCurrent() *animation.Animation {
	return e.Animation.Animation[e.GetState()]
}
func (e Actor) IterFrame() {
	e.GetCurrent().IterFrame()
}
func (e Actor) GetHealth() *component.Health {
	return e.Health
}
func (e Actor) GetActorState() *component.ActorState {
	return e.ActorState
}
