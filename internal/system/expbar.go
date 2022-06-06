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
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Expbar struct {
	jellyBeanQuery *query.Query
	query          *query.Query
}

func NewExpbar() *Expbar {
	return &Expbar{
		jellyBeanQuery: query.NewQuery(filter.Contains(component.JellyBeanTag)),
		query:          query.NewQuery(filter.Contains(entity.ExpBarTag)),
	}
}

func (ne *Expbar) Update(world donburi.World) {
	ne.query.EachEntity(world, func(entry *donburi.Entry) {
		exp := component.GetExp(entry)
		nextLevel := component.GetNextLevel(entry)

		if exp.CurrentExp >= exp.DesiredExp {

			ne.jellyBeanQuery.EachEntity(world, func(entry *donburi.Entry) {
				xp := (*component.XPData)(component.GetXP(entry))
				xp.Value = xp.Value / 2
			})

			exp.CurrentExp = 0
			nextLevel.CurrentLevel += 1
		}
	})
}

func (ne *Expbar) Draw(world donburi.World, screen *ebiten.Image) {
	ne.query.EachEntity(world, func(entry *donburi.Entry) {
		exp := component.GetExp(entry)
		position := component.GetPosition(entry)
		nextLevel := component.GetNextLevel(entry)

		ebitenutil.DrawRect(screen, position.X, position.Y, exp.DesiredExp*100, 16, colornames.Black)
		ebitenutil.DrawRect(screen, position.X, position.Y, exp.CurrentExp*100, 16, colornames.Blue700)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Lvl %d", nextLevel.CurrentLevel), (int(position.X)+int(exp.DesiredExp)*100)-50, int(position.Y))
	})
}
