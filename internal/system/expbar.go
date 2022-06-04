package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Expbar struct {
	query *query.Query
}

func NewExpbar() *Expbar {
	return &Expbar{
		query: query.NewQuery(filter.Contains(entity.ExpBarTag)),
	}
}

func (ne *Expbar) Update(world donburi.World) {
}

func (ne *Expbar) Draw(world donburi.World, screen *ebiten.Image) {
	ne.query.EachEntity(world, func(entry *donburi.Entry) {
		exp := component.GetExp(entry)
		position := component.GetPosition(entry)
		ebitenutil.DrawRect(screen, position.X, position.Y, exp.DesiredExp*100, 14, colornames.Black)
		ebitenutil.DrawRect(screen, position.X, position.Y, exp.CurrentExp*100, 14, colornames.Blue700)
	})
}
