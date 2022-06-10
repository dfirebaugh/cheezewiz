package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor interface {
	GetHealth() *component.Health
	GetState() *component.State
}
type Collidable interface {
	GetRigidBody() *component.RigidBody
	GetPosition() *component.Position
	GetState() *component.State
}
type Enemy interface {
	GetPosition() *component.Position
	GetHealth() *component.Health
	GetState() *component.State
	GetEnemyTag() ecs.Tag
}
type Player interface {
	GetHealth() *component.Health
	GetPosition() *component.Position
	GetInputDevice() input.PlayerInput
	GetState() *component.State
}
type Projectile interface {
	GetDirection() *component.Direction
	GetPosition() *component.Position
	GetProjectileTag() ecs.Tag
	GetHealth() *component.Health
}
type Animatable interface {
	GetFrame() *ebiten.Image
	GetPosition() *component.Position
	IterFrame()
	GetHealth() *component.Health
}
type ViewPort interface {
	ViewPort() ecs.Tag
	GetPosition() *component.Position
}
