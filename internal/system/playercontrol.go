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

		if controller.Device.IsUpPressed() {
			position.Y -= playerSpeed
		}

		if controller.Device.IsDownPressed() {
			position.Y += playerSpeed
		}

		if controller.Device.IsRightPressed() {
			position.X += playerSpeed
		}

		if controller.Device.IsLeftPressed() {
			position.X -= playerSpeed
		}
	})
}
