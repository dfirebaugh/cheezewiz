package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/yohamta/ganim8/v2"
)

type Renderable interface {
	Draw(screen *ebiten.Image)
}

type Render struct {
	count       int
	playerQuery *query.Query
}

func NewRender() *Render {
	return &Render{
		playerQuery: query.NewQuery(filter.Contains(entity.PlayerTag)),
	}
}

func (r *Render) Update(w donburi.World) {
	r.count++
	r.updatePlayer(w)
}

func (r Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.renderPlayer(w, screen)
}

func (r Render) updatePlayer(w donburi.World) {
	now := time.Now()

	r.playerQuery.EachEntity(w, func(entry *donburi.Entry) {
		animation := component.GetAnimation(entry)
		animation.Walk.Animation.Update(now.Sub(animation.Walk.PrevUpdateTime))
		animation.Walk.PrevUpdateTime = now
	})
}
func (r Render) renderPlayer(w donburi.World, screen *ebiten.Image) {
	r.playerQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		animation := component.GetAnimation(entry)
		op := ganim8.DrawOpts(position.X, position.Y)
		animation.Walk.Animation.Draw(screen, op)
	})
}
