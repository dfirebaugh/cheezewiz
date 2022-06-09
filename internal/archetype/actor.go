package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	*component.AnimationData
	*component.ActorStateData
	*component.PositionData
	*component.HealthAspect
}

func (a Actor) GetFrame() *ebiten.Image {
	return a.AnimationData.Animations[string(a.ActorStateData.GetCurrent())].GetFrame()
}
func (a Actor) GetPosition() *component.PositionData {
	return a.PositionData
}
func (a Actor) GetState() component.ActorStateType {
	return a.ActorStateData.GetCurrent()
}
func (a Actor) GetCurrent() *animation.Animation {
	return a.Animations[string(a.GetState())]
}
func (a Actor) IterFrame() {
	a.GetCurrent().IterFrame()
}
func (a Actor) GetHealth() *component.HealthAspect {
	return a.HealthAspect
}
