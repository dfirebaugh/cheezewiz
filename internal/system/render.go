package system

import (
	"cheezewiz/config"
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
	playerSlot         *query.Query
	tilemap_cache      *ebiten.Image
}

func NewRender() *Render {
	return &Render{
		playerQuery:        query.NewQuery(filter.Contains(entity.PlayerTag)),
		enemyQuery:         query.NewQuery(filter.Contains(entity.EnemyTag)),
		backgroundQuery:    query.NewQuery(filter.Contains(entity.BackgroundTag)),
		worldViewPortQuery: query.NewQuery(filter.Contains(entity.WorldViewPortTag)),
		playerSlot:         query.NewQuery(filter.Contains(entity.SlotTag)),
		tilemap_cache:      nil,
	}
}

func (r *Render) Update(w donburi.World) {
	r.count++
	r.updateEnemy(w)
	r.updatePlayer(w)
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.tileMap(w, screen)
	r.enemy(w, screen)
	r.player(w, screen)
	r.playerSlots(w, screen)
}

func (r *Render) updatePlayer(w donburi.World) {
	now := time.Now()

	r.playerQuery.EachEntity(w, func(entry *donburi.Entry) {
		animation := component.GetAnimation(entry)

		animation.Walk.Animation.Update(now.Sub(animation.Walk.PrevUpdateTime))
		animation.Hurt.Animation.Update(now.Sub(animation.Hurt.PrevUpdateTime))
		animation.Idle.Animation.Update(now.Sub(animation.Idle.PrevUpdateTime))
		animation.Walk.PrevUpdateTime = now
		animation.Hurt.PrevUpdateTime = now
		animation.Idle.PrevUpdateTime = now
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

func (r *Render) player(w donburi.World, screen *ebiten.Image) {
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
func (r Render) enemy(w donburi.World, screen *ebiten.Image) {
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

func (r *Render) tileMap(w donburi.World, screen *ebiten.Image) {

	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.backgroundQuery.EachEntity(w, func(entry *donburi.Entry) {

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0-worldViewLocationPos.X, 0-worldViewLocationPos.Y)

		if r.tilemap_cache == nil {
			println("Creating new tile cache")
			tiles := component.GetTileMap(entry)
			renderer, err := render.NewRenderer(tiles.Map)
			if err != nil {
				fmt.Printf("map unsupported for rendering: %s", err.Error())
				os.Exit(2)
			}
			if err = renderer.RenderVisibleLayers(); err != nil {
				fmt.Println(err)
				return
			}
			r.tilemap_cache = ebiten.NewImageFromImage(renderer.Result)
		}

		screen.DrawImage(r.tilemap_cache, op)
	})
}

func (r *Render) playerSlots(w donburi.World, screen *ebiten.Image) {
	padding := float64(15)
	r.playerSlot.EachEntity(w, func(entry *donburi.Entry) {
		sprite := component.GetSpriteSheet(entry)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(padding, float64(config.Get().Window.Width/config.Get().ScaleFactor)-480)
		screen.DrawImage(sprite.IMG, op)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(config.Get().Window.Height/config.Get().ScaleFactor)-500, float64(config.Get().Window.Width/config.Get().ScaleFactor)-80)
		screen.DrawImage(sprite.IMG, op)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(config.Get().Window.Height/config.Get().ScaleFactor)-142, float64(config.Get().Window.Width/config.Get().ScaleFactor)-80)
		screen.DrawImage(sprite.IMG, op)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(config.Get().Window.Height/config.Get().ScaleFactor)-142, float64(config.Get().Window.Width/config.Get().ScaleFactor)-480)
		screen.DrawImage(sprite.IMG, op)
	})
}
