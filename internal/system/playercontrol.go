package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type PlayerControl struct {
	query *query.Query
}

const playerSpeed = 1

func NewPlayerControl() *PlayerControl {
	return &PlayerControl{
		query: query.NewQuery(filter.Contains(entity.PlayerTag)),
	}
}

func (p PlayerControl) Update(w donburi.World) {
	p.query.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		controller := component.GetInputDevice(entry)
		state := component.GetActorState(entry)
		animation := component.GetAnimation(entry)

		isMovingRight := func () bool {
			return controller.Device.IsRightPressed() || controller.Device.IsRightJustPressed()
		}

		isMovingLeft := func () bool {
			return controller.Device.IsLeftPressed() || controller.Device.IsLeftJustPressed()
		}

		updatePlayerMovement := func (xPos float64, yPos float64) {
			position.Update(xPos, yPos)
			state.Set(component.WalkingState)
		}
		
		state.Reset()
		
		// Must check right and left first to have player facing in the correct direction
		if isMovingRight() {
			animation.Get(state.Current).Sprite.SetFlipH(true)

			if controller.Device.IsRightPressed() {
				updatePlayerMovement(position.X+playerSpeed, position.Y)
			}
		}

		if isMovingLeft() {
			animation.Get(state.Current).Sprite.SetFlipH(false)
			if controller.Device.IsLeftPressed() {
				updatePlayerMovement(position.X-playerSpeed, position.Y)
			}
		}

		if controller.Device.IsUpPressed() {
			updatePlayerMovement(position.X, position.Y-playerSpeed)
		}

		if controller.Device.IsDownPressed() {
			updatePlayerMovement(position.X, position.Y+playerSpeed)
		}
	})
}