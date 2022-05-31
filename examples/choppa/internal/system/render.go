package system

import (
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/examples/choppa/internal/entity"
	"image"

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

type Render struct {
	playerQuery     *query.Query
	projectileQuery *query.Query
	count           int
}

func NewRender() *Render {
	return &Render{
		playerQuery:     query.NewQuery(filter.Contains(entity.PlayerTag)),
		projectileQuery: query.NewQuery(filter.Contains(entity.ProjectileTag)),
	}
}

func (r *Render) Update(w donburi.World) {
	r.count++
}

func (r Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.renderPlayer(w, screen)
	r.renderProjectciles(w, screen)
}

func (r Render) renderPlayer(w donburi.World, screen *ebiten.Image) {
	r.playerQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		spriteSheet := component.GetSpriteSheet(entry)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)

		if r.count%10 == 0 {
			screen.DrawImage(spriteSheet.IMG.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), op)
			return
		}

		screen.DrawImage(spriteSheet.IMG.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image), op)
	})
}

func (r Render) renderProjectciles(w donburi.World, screen *ebiten.Image) {
	if r.projectileQuery.Count(w) == 0 {
		return
	}
	r.projectileQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)

		ebitenutil.DrawRect(
			screen,
			position.X, position.Y+32/2,
			100, 1,
			colornames.Tomato,
		)
	})
}
