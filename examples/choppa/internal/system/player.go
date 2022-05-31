package system

import (
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/examples/choppa/internal/entity"
	"cheezewiz/internal/input"

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

		k := input.Keyboard{}

		if k.IsUpPressed() {
			position.Y -= playerSpeed
		}

		if k.IsDownPressed() {
			position.Y += playerSpeed
		}

		if k.IsRightJustPressed() {
			// p.flipRight(entry)
		}
		if k.IsRightPressed() {
			position.X += playerSpeed
		}

		if k.IsLeftJustPressed() {
			// p.flipLeft(entry)
		}
		if k.IsLeftPressed() {
			position.X -= playerSpeed
		}

		if k.IsSpaceJustPressed() {
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
