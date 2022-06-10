package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	*component.Animation
	*component.State
	*component.Position
	*component.RigidBody
	*component.Direction
	*component.Health
	ProjectileTag ecs.Tag
}

func (p Projectile) GetProjectileTag() ecs.Tag {
	return p.ProjectileTag
}
func (p Projectile) GetRigidBody() *component.RigidBody {
	return p.RigidBody
}
func (p Projectile) GetFrame() *ebiten.Image {
	return p.Animation.Animation[p.State.GetCurrent()].GetFrame()
}
func (p Projectile) GetPosition() *component.Position {
	return p.Position
}
func (p Projectile) GetCurrentState() component.StateType {
	return p.State.GetCurrent()
}
func (p Projectile) GetCurrent() *animation.Animation {
	return p.Animation.Animation[p.GetCurrentState()]
}
func (p Projectile) IterFrame() {
	p.GetCurrent().IterFrame()
}
func (p Projectile) GetState() *component.State {
	return p.State
}
func (p Projectile) GetDirection() *component.Direction {
	return p.Direction
}
func (p Projectile) GetHealth() *component.Health {
	return p.Health
}
