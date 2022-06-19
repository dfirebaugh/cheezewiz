package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"
	"cheezewiz/pkg/gamemath"
)

type Controller struct {
	w world.World
}

const playerSpeed = 1

func MakePlayerControl(w world.World) Controller {
	return Controller{
		w: w,
	}
}

func (c Controller) Update() {
	query.Each(c.w, filter.GetPlayers, func(handle world.EntityHandle) {
		c.controllable(handle)
	})
}

func (c Controller) controllable(handle world.EntityHandle) {
	p := c.w.GetEntity(handle)

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
		if controller.IsRightPressed() {
			updatePlayerMovement(position.X+playerSpeed, position.Y)
		}
	}

	if isMovingLeft() {
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
	}

	// health := component.GetHealth(entry)
	// if health.HP <= 0 {
	// 	state.Set(component.DeathState)
	// }
}

func (c Controller) setTilePosition(position component.Position) {

}
