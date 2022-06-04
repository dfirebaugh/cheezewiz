package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

var tpsCount = 0

type Timer struct {
	query *query.Query
}

func NewTimer() *Timer {
	return &Timer{
		query: query.NewQuery(filter.Contains(entity.TimerTag)),
	}
}

func (t *Timer) Update(world donburi.World) {
	t.query.EachEntity(world, func(entry *donburi.Entry) {
		countdown := component.GetCountdown(entry)
		tpsCount++
		if tpsCount%60 == 0 && countdown.CountDownInSec > 0 {
			countdown.CountDownInSec--
		}
	})
}

func (t *Timer) Draw(world donburi.World, screen *ebiten.Image) {
	t.query.EachEntity(world, func(entry *donburi.Entry) {
		countdown := component.GetCountdown(entry)
		position := component.GetPosition(entry)
		ebitenutil.DebugPrintAt(screen, displayTotalCountdown(countdown.CountDownInSec), int(position.X), int(position.Y))
	})
}

func displayTotalCountdown(countdown uint32) string {
	var remainingSecs = countdown % 60

	displayRemainingSecs := func() string {
		if remainingSecs <= 9 {
			return fmt.Sprintf("0%d", remainingSecs)
		}

		return fmt.Sprintf("%d", remainingSecs)
	}
	return fmt.Sprintf("%d:%s", int(countdown/60), displayRemainingSecs())
}
