package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor interface {
	GetPosition() *component.PositionData
	GetHealth() *component.HealthAspect
}

type ActorArchetype struct {
	*component.AnimationData
	*component.ActorStateData
	*component.PositionData
	*component.HealthAspect
}

func (a ActorArchetype) GetFrame() *ebiten.Image {
	return a.AnimationData.Animations[string(a.ActorStateData.GetCurrent())].GetFrame()
}
func (a ActorArchetype) GetPosition() *component.PositionData {
	return a.PositionData
}
func (a ActorArchetype) GetState() component.ActorStateType {
	return a.ActorStateData.GetCurrent()
}
func (a ActorArchetype) GetCurrent() *animation.Animation {
	return a.Animations[string(a.GetState())]
}
func (a ActorArchetype) IterFrame() {
	a.GetCurrent().IterFrame()
}
func (a ActorArchetype) GetHealth() *component.HealthAspect {
	return a.HealthAspect
}
