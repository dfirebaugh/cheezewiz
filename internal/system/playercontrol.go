package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/attacks"
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/tag"
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"
	"cheezewiz/pkg/gamemath"
	"cheezewiz/pkg/taskrunner"
	"time"
)

type PlayerController struct{}

const playerSpeed = 1

func (c PlayerController) addPlayer() {
	entity.MakeWithTags(world.Instance, "entities/cheezewiz.entity.json",
		float64(config.Get().Window.Width/config.Get().ScaleFactor/2),
		float64(config.Get().Window.Height/config.Get().ScaleFactor/2), []tag.Tag{tag.Player, tag.Animatable, tag.Collidable})

	taskrunner.Add(time.Millisecond*800, attacks.CheeseMissile())
}

func (c PlayerController) Update() {
	if query.Count(world.Instance, filter.GetPlayers) == 0 {
		c.addPlayer()
	}
	query.Each(world.Instance, filter.GetPlayers, func(handle world.EntityHandle) {
		c.controllable(handle)
	})
}

func (c PlayerController) controllable(handle world.EntityHandle) {
	p := world.Instance.GetEntity(handle)

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
