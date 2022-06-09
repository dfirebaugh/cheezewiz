package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyTag struct{}
type Enemy interface {
	GetPosition() *component.Position
	GetHealth() *component.Health
	GetEnemyTag() EnemyTag
}

type EnemyArchetype struct {
	*component.Animation
	*component.ActorState
	*component.Position
	*component.Health
	EnemyTag
}

func (e EnemyArchetype) GetEnemyTag() EnemyTag {
	return e.EnemyTag
}
func (e EnemyArchetype) GetFrame() *ebiten.Image {
	return e.Animation.Animation[string(e.ActorState.GetCurrent())].GetFrame()
}
func (e EnemyArchetype) GetPosition() *component.Position {
	return e.Position
}
func (e EnemyArchetype) GetState() component.ActorStateType {
	return e.ActorState.GetCurrent()
}
func (e EnemyArchetype) GetCurrent() *animation.Animation {
	return e.Animation.Animation[string(e.GetState())]
}
func (e EnemyArchetype) IterFrame() {
	e.GetCurrent().IterFrame()
}
func (e EnemyArchetype) GetHealth() *component.Health {
	return e.Health
}
