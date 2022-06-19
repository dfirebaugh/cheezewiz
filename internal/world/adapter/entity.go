package adapter

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/internal/tag"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetHealth() *component.Health
	GetState() *component.State
	GetDirection() *component.Direction
	GetPosition() *component.Position
	GetRigidBody() *component.RigidBody
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions)
	DebugDraw(screen *ebiten.Image)
	IterFrame()
	AddTag(t tag.Tag)
	HasTag(t tag.Tag) bool
	GetInputDevice() input.PlayerInput
}
