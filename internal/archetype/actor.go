package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor interface {
	GetPosition() *component.Position
	GetHealth() *component.Health
	GetRigidBody() *component.RigidBody
}

type ActorArchetype struct {
	*component.Animation
	*component.ActorState
	*component.Position
	*component.Health
	*component.Direction
	*component.RigidBody
}

func (a ActorArchetype) GetFrame() *ebiten.Image {
	return a.Animation.Animation[string(a.ActorState.GetCurrent())].GetFrame()
}
func (a ActorArchetype) GetPosition() *component.Position {
	return a.Position
}
func (a ActorArchetype) GetState() component.ActorStateType {
	return a.ActorState.GetCurrent()
}
func (a ActorArchetype) GetCurrent() *animation.Animation {
	return a.Animation.Animation[string(a.GetState())]
}
func (a ActorArchetype) IterFrame() {
	a.GetCurrent().IterFrame()
}
func (a ActorArchetype) GetHealth() *component.Health {
	return a.Health
}
func (a ActorArchetype) GetDirection() *component.Direction {
	return a.Direction
}
func (a ActorArchetype) GetRigidBody() *component.RigidBody {
	return a.RigidBody
}
