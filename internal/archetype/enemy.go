package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyTag struct{}
type Enemy interface {
	GetPosition() *component.PositionData
	GetHealth() *component.HealthAspect
	GetEnemyTag() EnemyTag
}

type EnemyArchetype struct {
	*component.AnimationData
	*component.ActorStateData
	*component.PositionData
	*component.HealthAspect
	EnemyTag
}

func (e EnemyArchetype) GetEnemyTag() EnemyTag {
	return e.EnemyTag
}
func (e EnemyArchetype) GetFrame() *ebiten.Image {
	return e.AnimationData.Animations[string(e.ActorStateData.GetCurrent())].GetFrame()
}
func (e EnemyArchetype) GetPosition() *component.PositionData {
	return e.PositionData
}
func (e EnemyArchetype) GetState() component.ActorStateType {
	return e.ActorStateData.GetCurrent()
}
func (e EnemyArchetype) GetCurrent() *animation.Animation {
	return e.Animations[string(e.GetState())]
}
func (e EnemyArchetype) IterFrame() {
	e.GetCurrent().IterFrame()
}
func (e EnemyArchetype) GetHealth() *component.HealthAspect {
	return e.HealthAspect
}
