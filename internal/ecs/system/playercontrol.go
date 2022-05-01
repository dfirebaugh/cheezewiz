package system

import (
	"cheezewiz/internal/ecs/component"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type controller interface {
	IsUpPressed() bool
	IsDownPressed() bool
	IsLeftPressed() bool
	IsRightPressed() bool
}

type playerControl struct {
	Position   *component.Position
	Velocity   *component.Velocity
	Asset      *component.Asset
	controller controller
}

func NewPlayerControl(controller controller) gohan.System {
	return &playerControl{controller: controller}
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
	if i.controller.IsUpPressed() {
		i.accelerateUp()
	}
	if i.controller.IsLeftPressed() {
		i.accelerateLeft()
	}
	if i.controller.IsDownPressed() {
		i.accelerateDown()
	}
	if i.controller.IsRightPressed() {
		i.accelerateRight()
	}
}
