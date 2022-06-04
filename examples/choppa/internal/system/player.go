package system

import (
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/examples/choppa/internal/entity"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Player struct {
	query *query.Query
}

const playerSpeed = 1

func NewPlayer() *Player {
	return &Player{
		query: query.NewQuery(filter.Contains(entity.PlayerTag)),
	}
}

func (p Player) Update(w donburi.World) {
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

		if controller.Device.IsPrimaryAtkJustPressed() {
			entity.MakeProjectile(w, position)
		}
	})
}

// func (p Player) flipRight(entry *donburi.Entry) {
// 	direction := component.GetDirection(entry)
// 	if !direction.IsRight {
// 		direction.IsRight = true
// 	}
// }

// func (p Player) flipLeft(entry *donburi.Entry) {
// 	direction := component.GetDirection(entry)
// 	if direction.IsRight {
// 		direction.IsRight = false
// 	}
// }
