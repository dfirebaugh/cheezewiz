package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/yohamta/ganim8/v2"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Renderable interface {
	Draw(screen *ebiten.Image)
}

type Render struct {
	count              int
	playerQuery        *query.Query
	enemyQuery         *query.Query
	backgroundQuery    *query.Query
	worldViewPortQuery *query.Query
}

func NewRender() *Render {
	return &Render{
		playerQuery:        query.NewQuery(filter.Contains(entity.PlayerTag)),
		enemyQuery:         query.NewQuery(filter.Contains(entity.EnemyTag)),
		backgroundQuery:    query.NewQuery(filter.Contains(entity.BackgroundTag)),
		worldViewPortQuery: query.NewQuery(filter.Contains(entity.WorldViewPortTag)),
	}
}

func (r *Render) Update(w donburi.World) {
	r.count++
	r.updateEnemy(w)
	r.updatePlayer(w)

}

func (r Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.renderTileMap(w, screen)
	r.renderEnemy(w, screen)
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

func (r Render) updateEnemy(w donburi.World) {
	now := time.Now()

	r.enemyQuery.EachEntity(w, func(entry *donburi.Entry) {
		animation := component.GetAnimation(entry)
		animation.Walk.Animation.Update(now.Sub(animation.Walk.PrevUpdateTime))
		animation.Walk.PrevUpdateTime = now
	})
}

func (r Render) renderPlayer(w donburi.World, screen *ebiten.Image) {
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.playerQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		health := component.GetHealth(entry)
		animation := component.GetAnimation(entry)
		direction := component.GetDirection(entry)
		state := component.GetPlayerState(entry)

		op := ganim8.DrawOpts(position.X-worldViewLocationPos.X, position.Y-worldViewLocationPos.Y)

		if direction.IsRight {
			op.SetScale(-1, 1)
			op.SetPos(position.X+32-worldViewLocationPos.X, position.Y-worldViewLocationPos.Y)
		}

		switch state.Current {
		case component.WalkingState:
			animation.Walk.Animation.Draw(screen, op)
		case component.AttackingState:
		case component.HurtState:
			animation.Hurt.Animation.Draw(screen, op)
		case component.DeathState:
		default:
			animation.Idle.Animation.Draw(screen, op)
		}

		state.ResetState()

		ebitenutil.DrawRect(screen, position.X-worldViewLocationPos.X, position.Y+35-worldViewLocationPos.Y, health.MAXHP/3, 3, colornames.Grey100)
		ebitenutil.DrawRect(screen, position.X-worldViewLocationPos.X, position.Y+35-worldViewLocationPos.Y, health.HP/3, 3, colornames.Red600)
	})
}
func (r Render) renderEnemy(w donburi.World, screen *ebiten.Image) {
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.enemyQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		animation := component.GetAnimation(entry)
		op := ganim8.DrawOpts(position.X-worldViewLocationPos.X, position.Y-worldViewLocationPos.Y)
		// op.SetScale(-1, 0)
		animation.Walk.Animation.Draw(screen, op)
	})
}

func (r Render) renderTileMap(w donburi.World, screen *ebiten.Image) {
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.backgroundQuery.EachEntity(w, func(entry *donburi.Entry) {
		tiles := component.GetTileMap(entry)

		renderer, err := render.NewRenderer(tiles.Map)
		if err != nil {
			fmt.Printf("map unsupported for rendering: %s", err.Error())
			os.Exit(2)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0-worldViewLocationPos.X, 0-worldViewLocationPos.Y)
		if err = renderer.RenderVisibleLayers(); err != nil {
			fmt.Println(err)
			return
		}
		screen.DrawImage(ebiten.NewImageFromImage(renderer.Result), op)
	})
}
