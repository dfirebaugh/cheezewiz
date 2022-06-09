package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*component.Animation
	*component.ActorState
	*component.InputDevice
	*component.Position
	*component.Health
	*component.RigidBody
	PlayerTag ecs.Tag
}

func (p Player) GetPlayerTag() ecs.Tag {
	return p.PlayerTag
}
func (p Player) GetRigidBody() *component.RigidBody {
	return p.RigidBody
}
func (p Player) GetInputDevice() input.PlayerInput {
	return p.InputDevice.Device
}
func (p Player) GetFrame() *ebiten.Image {
	return p.Animation.Animation[string(p.ActorState.GetCurrent())].GetFrame()
}
func (p Player) GetPosition() *component.Position {
	return p.Position
}
func (p Player) GetState() component.ActorStateType {
	return p.ActorState.GetCurrent()
}
func (p Player) GetCurrent() *animation.Animation {
	return p.Animation.Animation[string(p.GetState())]
}
func (p Player) IterFrame() {
	p.GetCurrent().IterFrame()
}
func (p Player) GetHealth() *component.Health {
	return p.Health
}
func (p Player) GetActorState() *component.ActorState {
	return p.ActorState
}
