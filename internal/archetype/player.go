package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player interface {
	GetPosition() *component.Position
	GetHealth() *component.Health
	GetInputDevice() input.PlayerInput
	GetActorState() *component.ActorState
}

type PlayerArchetype struct {
	*component.Animation
	*component.ActorState
	*component.InputDevice
	*component.Position
	*component.Health
}

func (p PlayerArchetype) GetInputDevice() input.PlayerInput {
	return p.InputDevice.Device
}
func (p PlayerArchetype) GetFrame() *ebiten.Image {
	return p.Animation.Animation[string(p.ActorState.GetCurrent())].GetFrame()
}
func (p PlayerArchetype) GetPosition() *component.Position {
	return p.Position
}
func (p PlayerArchetype) GetState() component.ActorStateType {
	return p.ActorState.GetCurrent()
}
func (p PlayerArchetype) GetCurrent() *animation.Animation {
	return p.Animation.Animation[string(p.GetState())]
}
func (p PlayerArchetype) IterFrame() {
	p.GetCurrent().IterFrame()
}
func (p PlayerArchetype) GetHealth() *component.Health {
	return p.Health
}
func (p PlayerArchetype) GetActorState() *component.ActorState {
	return p.ActorState
}
