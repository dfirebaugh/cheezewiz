package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player interface {
	GetPosition() *component.PositionData
	GetHealth() *component.HealthAspect
	GetInputDevice() input.PlayerInput
}

type PlayerArchetype struct {
	*component.AnimationData
	*component.ActorStateData
	*component.InputDeviceData
	*component.PositionData
	*component.HealthAspect
}

func (p PlayerArchetype) GetInputDevice() input.PlayerInput {
	return p.InputDeviceData.Device
}
func (p PlayerArchetype) GetFrame() *ebiten.Image {
	return p.AnimationData.Animations[string(p.ActorStateData.GetCurrent())].GetFrame()
}
func (p PlayerArchetype) GetPosition() *component.PositionData {
	return p.PositionData
}
func (p PlayerArchetype) GetState() component.ActorStateType {
	return p.ActorStateData.GetCurrent()
}
func (p PlayerArchetype) GetCurrent() *animation.Animation {
	return p.Animations[string(p.GetState())]
}
func (p PlayerArchetype) IterFrame() {
	p.GetCurrent().IterFrame()
}
func (p PlayerArchetype) GetHealth() *component.HealthAspect {
	return p.HealthAspect
}
