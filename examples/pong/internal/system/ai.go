package system

import (
	"cheezewiz/examples/pong/internal/component"
	"cheezewiz/examples/pong/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type AI struct {
	query     *query.Query
	ballQuery *query.Query
}

func NewAI() *AI {
	return &AI{
		query:     query.NewQuery(filter.Contains(entity.EnemyTag)),
		ballQuery: query.NewQuery(filter.Contains(entity.BallTag)),
	}
}

func (a *AI) Update(w donburi.World) {
	_, screenHeight := ebiten.WindowSize()
	a.ballQuery.EachEntity(w, func(ball *donburi.Entry) {
		ballPos := component.GetPosition(ball)

		a.query.EachEntity(w, func(entry *donburi.Entry) {
			pos := component.GetPosition(entry)
			rect := component.GetRect(entry)

			if int(ballPos.Y) < int(pos.Y+(rect.Height/2)) {
				if pos.Y <= 0 {
					return
				}
				pos.Y += -3
			}
			if int(ballPos.Y) > int(pos.Y+(rect.Height/2)) {
				if int(pos.Y) >= screenHeight {
					return
				}
				pos.Y += 3
			}
		})
	})
}
