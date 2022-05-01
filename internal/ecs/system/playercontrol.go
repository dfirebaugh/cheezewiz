package system

import (
	"cheezewiz/internal/ecs/component"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type playerControl struct {
	Position   *component.Position
	Velocity   *component.Velocity
	Animation  *component.Animation
	Controller *component.Controller
}

func NewPlayerControl() gohan.System {
	return &playerControl{}
}

const (
	topSpeed     = 3.3
	acceleration = 0.22
	deceleration = 8.88
)

func (i *playerControl) Update(entity gohan.Entity) error {
	i.handleInput()
	i.handleDeceleration()
	return nil
}

func (i *playerControl) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}

// handleDeceleration asymptotes toward zero
func (i *playerControl) handleDeceleration() {
	// add a limit so we don't forever go toward zero
	if i.Velocity.Y >= 0.002 || i.Velocity.Y <= -0.002 {
		i.Velocity.Y = i.Velocity.Y - (i.Velocity.Y / deceleration)
	}

	if i.Velocity.X >= 0.002 || i.Velocity.X <= -0.002 {
		i.Velocity.X = i.Velocity.X - (i.Velocity.X / deceleration)
	}
}

func (i *playerControl) accelerateUp() {
	if i.Velocity.Y <= -topSpeed {
		return
	}
	i.Velocity.Y -= acceleration
	i.Velocity.Y = -topSpeed
}

func (i *playerControl) accelerateLeft() {
	if i.Velocity.X <= -topSpeed {
		return
	}
	i.Velocity.X -= acceleration
	i.Velocity.X = -topSpeed
}

func (i *playerControl) accelerateDown() {
	if i.Velocity.Y >= topSpeed {
		return
	}
	i.Velocity.Y += acceleration
	i.Velocity.Y = topSpeed
}

func (i *playerControl) accelerateRight() {
	if i.Velocity.X >= topSpeed {
		return
	}
	i.Velocity.X += acceleration
	i.Velocity.X = topSpeed
}

func (i *playerControl) handleInput() {
	if i.Controller.Controller.IsUpPressed() {
		i.accelerateUp()
	}
	if i.Controller.Controller.IsLeftPressed() {
		i.accelerateLeft()
	}
	if i.Controller.Controller.IsDownPressed() {
		i.accelerateDown()
	}
	if i.Controller.Controller.IsRightPressed() {
		i.accelerateRight()
	}
}
