package system

import (
	"cheezewiz/examples/pongv2/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/image/colornames"
)

type Renderable interface {
	Draw(screen *ebiten.Image)
}

// When you need many sets of components
// in one system, you can use the views
type Render struct {
	query     *query.Query
	ballQuery *query.Query
}

func NewRender() *Render {
	// w.AddSystems(&Render{})
	return &Render{
		query: query.NewQuery(filter.Contains(
			component.Position,
			component.Rect,
		)),
		ballQuery: query.NewQuery(filter.Contains(
			component.Position,
			component.Radius,
			component.Velocity,
		))}
}

// Render one frame
func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	// But choose the right entities yourself
	r.renderBall(w, screen)
	r.renderBorders(w, screen)
}

func (r Render) renderBall(w donburi.World, screen *ebiten.Image) {
	r.ballQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		radius := component.GetRadius(entry)

		ebitenutil.DrawRect(
			screen,
			position.X-radius.Value/2, position.Y-radius.Value/2,
			radius.Value, radius.Value,
			colornames.Aliceblue,
		)
	})
}

func (r Render) renderBorders(w donburi.World, screen *ebiten.Image) {
	r.query.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		rect := component.GetRect(entry)

		ebitenutil.DrawRect(
			screen,
			position.X, position.Y,
			rect.Width, rect.Height,
			colornames.Aliceblue,
		)
	})
}
