package system

import (
	"cheezewiz/examples/pongv2/internal/component"
	"cheezewiz/examples/pongv2/internal/entity"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Score struct {
	Player    uint
	Enemy     uint
	ballQuery *query.Query
}

func NewScore() *Score {
	return &Score{
		Player:    0,
		Enemy:     0,
		ballQuery: query.NewQuery(filter.Contains(entity.BallTag)),
	}
}

func (s *Score) Update(w donburi.World) {
	screenWidth, _ := ebiten.WindowSize()
	s.ballQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		if position.X < 0 {
			s.Enemy++
			s.ResetBall(entry)
		}
		if int(position.X) > screenWidth {
			s.Player++
			s.ResetBall(entry)
		}
	})
}

func (s *Score) Draw(w donburi.World, screen *ebiten.Image) {
	screenWidth, _ := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d:%d", s.Player, s.Enemy), screenWidth/2, 15)
}

func (s *Score) ResetBall(entry *donburi.Entry) {
	screenWidth, screenHeight := ebiten.WindowSize()
	position := component.GetPosition(entry)
	velocity := component.GetVelocity(entry)

	position.X = float64(screenWidth) / 2
	position.Y = float64(screenHeight) / 2
	velocity.L = 2
	velocity.M = 2
}
