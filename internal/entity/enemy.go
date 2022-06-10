package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	*component.Animation
	*component.ActorState
	*component.Position
	*component.Health
	*component.RigidBody
	EnemyTag ecs.Tag
}

func (e Enemy) GetEnemyTag() ecs.Tag {
	return e.EnemyTag
}
func (e Enemy) GetRigidBody() *component.RigidBody {
	return e.RigidBody
}
func (e Enemy) GetFrame() *ebiten.Image {
	current := e.ActorState.GetCurrent()

	return e.Animation.Animation[current].GetFrame()
}
func (e Enemy) GetPosition() *component.Position {
	return e.Position
}
func (e Enemy) GetState() component.ActorStateType {
	return e.ActorState.GetCurrent()
}
func (e Enemy) GetCurrent() *animation.Animation {
	return e.Animation.Animation[e.GetState()]
}
func (e Enemy) IterFrame() {
	e.GetCurrent().IterFrame()
}
func (e Enemy) GetHealth() *component.Health {
	return e.Health
}
func (e Enemy) GetActorState() *component.ActorState {
	return e.ActorState
}
