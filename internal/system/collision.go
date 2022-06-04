package system

import (
	"cheezewiz/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Collision struct{}

func NewCollision() *Collision {
	return &Collision{}
}

func (c *Collision) Update(w donburi.World) {
}

func (c *Collision) Draw(w donburi.World, screen *ebiten.Image) {
}

func (c *Collision) IsCollide(a component.RigidBodyData, b component.RigidBodyData) bool {
	return a.Y < b.Y+b.H && a.Y+a.H > b.Y && a.X < b.X+b.W && a.X+a.W > b.X
}
