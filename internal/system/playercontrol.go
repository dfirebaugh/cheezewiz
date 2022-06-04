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
		direction := component.GetDirection(entry)
		state := component.GetPlayerState(entry)

		if controller.Device.IsUpPressed() {
			position.Y -= playerSpeed
			state.Current = component.WalkingState
		}

		if controller.Device.IsDownPressed() {
			position.Y += playerSpeed
			state.Current = component.WalkingState
		}

		if controller.Device.IsRightPressed() {
			position.X += playerSpeed
			direction.IsRight = true
			state.Current = component.WalkingState
		}

		if controller.Device.IsLeftPressed() {
			position.X -= playerSpeed
			direction.IsRight = false
			state.Current = component.WalkingState
		}

		if controller.Device.IsPrimaryAtkJustPressed() {
			state.Current = component.HurtState
		}
	})
}
