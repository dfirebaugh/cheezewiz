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
		state := component.GetActorState(entry)

		state.Reset()

		if controller.Device.IsUpPressed() {
			position.Update(position.X, position.Y-playerSpeed)
			state.Set(component.WalkingState)
		}

		if controller.Device.IsDownPressed() {
			position.Update(position.X, position.Y+playerSpeed)
			state.Set(component.WalkingState)
		}

		if controller.Device.IsRightPressed() {
			position.Update(position.X+playerSpeed, position.Y)
			direction.IsRight = true
			state.Set(component.WalkingState)
		}

		if controller.Device.IsLeftPressed() {
			position.Update(position.X-playerSpeed, position.Y)
			direction.IsRight = false
			state.Set(component.WalkingState)
		}

		if controller.Device.IsPrimaryAtkJustPressed() {
		}
	})
}
