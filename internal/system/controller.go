package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/ecs"
)

type Controllable interface {
	GetInputDevice() input.PlayerInput
	GetPosition() *component.PositionData
}

type Controller struct {
	World ecs.World
}

func (c Controller) Update() {
	for _, player := range ecs.FilterBy[Controllable](c.World) {
		c.controllable(player)
	}
}

func (c Controller) controllable(e Controllable) {
	controller := e.GetInputDevice()
	position := e.GetPosition()
	// state := e.GetState()
	// animation := component.GetAnimation(entry)
	// direction := component.GetDirection(entry)

	isMovingRight := func() bool {
		return controller.IsRightPressed() || controller.IsRightJustPressed()
	}

	isMovingLeft := func() bool {
		return controller.IsLeftPressed() || controller.IsLeftJustPressed()
	}

	updatePlayerMovement := func(xPos float64, yPos float64) {
		position.Update(xPos, yPos)
		// state.Set(component.WalkingState)
	}

	// state.Reset()

	// Must check right and left first to have player facing in the correct direction
	if isMovingRight() {
		// animation.Get(state.Current).Sprite.SetFlipH(true)
		// direction.IsRight = true

		if controller.IsRightPressed() {
			updatePlayerMovement(position.X+playerSpeed, position.Y)
		}
	}

	if isMovingLeft() {
		// direction.IsRight = false
		// animation.Get(state.Current).Sprite.SetFlipH(false)
		if controller.IsLeftPressed() {
			updatePlayerMovement(position.X-playerSpeed, position.Y)
		}
	}

	if controller.IsUpPressed() {
		updatePlayerMovement(position.X, position.Y-playerSpeed)
	}

	if controller.IsDownPressed() {
		updatePlayerMovement(position.X, position.Y+playerSpeed)
		// state.Set(component.IdleState)
	}

	// health := component.GetHealth(entry)
	// if health.HP <= 0 {
	// 	state.Set(component.DeathState)
	// }
}
