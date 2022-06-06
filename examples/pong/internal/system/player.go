package system

import (
	"cheezewiz/examples/pong/internal/component"
	"cheezewiz/examples/pong/internal/entity"
	"cheezewiz/internal/input"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Player struct {
	query *query.Query
}

func NewPlayer() *Player {
	return &Player{
		query: query.NewQuery(filter.Contains(entity.PlayerTag)),
	}
}

func (p *Player) Update(w donburi.World) {
	p.query.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)

		k := input.Keyboard{}

		if k.IsUpPressed() {
			position.Y -= 4
		}

		if k.IsDownPressed() {
			position.Y += 4
		}
	})
}
