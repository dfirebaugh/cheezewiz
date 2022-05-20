package system

import (
	"cheezewiz/internal/ecs/component"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	topSpeed     = 3.3
	acceleration = 0.22
	deceleration = 8.88
)

type Controllable interface {
	GetPosition() *component.Position
	GetVelocity() *component.Velocity
	GetAnimation() *component.Animation
	GetController() *component.Controller
}

type Control struct {
	Level *component.Level
}

type ControllableEntity struct {
	Controllable
}

func (c *Control) AttachLevel(lvl *component.Level) {
	c.Level = lvl
}

func (c Control) Update() {
	for _, id := range c.Level.Entities {
		if _, ok := c.Level.EntityMap[id].(Controllable); !ok {
			continue
		}
		ctrl := c.Level.EntityMap[id].(Controllable)
		e := &ControllableEntity{
			Controllable: ctrl,
		}

		e.handleInput()
		e.handleDeceleration()
	}
}
func (c Control) Render(screen *ebiten.Image) {}

// handleDeceleration asymptotes toward zero
func (e *ControllableEntity) handleDeceleration() {
	// add a limit so we don't forever go toward zero
	if e.GetVelocity().Y >= 0.002 || e.GetVelocity().Y <= -0.002 {
		e.GetVelocity().Y = e.GetVelocity().Y - (e.GetVelocity().Y / deceleration)
	}

	if e.GetVelocity().X >= 0.002 || e.GetVelocity().X <= -0.002 {
		e.GetVelocity().X = e.GetVelocity().X - (e.GetVelocity().X / deceleration)
	}
}

func (e *ControllableEntity) accelerateUp() {
	if e.GetVelocity().Y <= -topSpeed {
		return
	}
	e.GetVelocity().Y -= acceleration
	e.GetVelocity().Y = -topSpeed
}

func (e *ControllableEntity) accelerateLeft() {
	if e.GetVelocity().X <= -topSpeed {
		return
	}
	e.GetVelocity().X -= acceleration
	e.GetVelocity().X = -topSpeed
}

func (e *ControllableEntity) accelerateDown() {
	if e.GetVelocity().Y >= topSpeed {
		return
	}
	e.GetVelocity().Y += acceleration
	e.GetVelocity().Y = topSpeed
}

func (e *ControllableEntity) accelerateRight() {
	if e.GetVelocity().X >= topSpeed {
		return
	}
	e.GetVelocity().X += acceleration
	e.GetVelocity().X = topSpeed
}

func (e *ControllableEntity) handleInput() {
	if e.GetController().Controller.IsUpPressed() {
		e.accelerateUp()
	}
	if e.GetController().Controller.IsLeftPressed() {
		e.accelerateLeft()
	}
	if e.GetController().Controller.IsDownPressed() {
		e.accelerateDown()
	}
	if e.GetController().Controller.IsRightPressed() {
		e.accelerateRight()
	}
}
