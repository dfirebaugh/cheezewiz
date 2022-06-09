package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerFilter = func(entity ecs.Entity) bool {
	if _, ok := entity.(Player); ok {
		return true
	}
	return false
}

type Player struct {
	*component.AnimationData
	*component.ActorStateData
	*component.InputDeviceData
	*component.PositionData
	*component.HealthAspect
}

func (p Player) GetInputDevice() input.PlayerInput {
	return p.InputDeviceData.Device
}
func (p Player) GetFrame() *ebiten.Image {
	return p.AnimationData.Animations[string(p.ActorStateData.GetCurrent())].GetFrame()
}
func (p Player) GetPosition() *component.PositionData {
	return p.PositionData
}
func (p Player) GetState() component.ActorStateType {
	return p.ActorStateData.GetCurrent()
}
func (p Player) GetCurrent() *animation.Animation {
	return p.Animations[string(p.GetState())]
}
func (p Player) IterFrame() {
	p.GetCurrent().IterFrame()
}
func (p Player) GetHealth() *component.HealthAspect {
	return p.HealthAspect
}
