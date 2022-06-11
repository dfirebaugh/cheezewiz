package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/internal/ecs/adapter"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/gamemath"
)

type Controller struct {
	ecs   adapter.Adapter
	Level archetype.Level
}

const playerSpeed = 1

func MakePlayerControl(ecs adapter.Adapter, level archetype.Level) Controller {
	return Controller{
		ecs:   ecs,
		Level: level,
	}
}

func (c Controller) Update() {
	playerHandles, _ := c.ecs.GetPlayerHandles()
	for _, playerHandle := range playerHandles {
		c.controllable(playerHandle)
	}
}

func (c Controller) controllable(playerHandle ecs.EntityHandle) {
	p, ok := c.ecs.GetEntity(playerHandle).(archetype.Player)
	if !ok {
		return
	}
	controller := p.GetInputDevice()
	position := p.GetPosition()
	health := p.GetHealth()
	state := p.GetState()
	if health.Current == 0 {
		state.Set(component.DeathState)
		return
	}

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

	if controller.IsPrimaryAtkJustPressed() {
		println(gamemath.Vector([]float64{position.X, position.Y}).ToTileCoord(float64(config.Get().TileSize)).ToString())
		// move the player from one tile to another
		// c.level.MoveActor(position, position, p, p)
		println(c.Level.ToString())
	}

	// health := component.GetHealth(entry)
	// if health.HP <= 0 {
	// 	state.Set(component.DeathState)
	// }
}

func (c Controller) setTilePosition(position component.Position) {

}
