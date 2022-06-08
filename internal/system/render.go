package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/tag"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled/render"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Renderable interface {
	Draw(screen *ebiten.Image)
}

type Render struct {
	count                int
	animatableActorQuery *query.Query
	rigidBodyQuery       *query.Query
	backgroundQuery      *query.Query
	worldViewPortQuery   *query.Query
	playerSlot           *query.Query
	jellyBeanQuery       *query.Query
	damageLabelQuery     *query.Query
	positionQuery        *query.Query
	tilemap_cache        *ebiten.Image
}

func NewRender() *Render {
	return &Render{
		animatableActorQuery: query.NewQuery(filter.Contains(component.Animation, component.ActorState)),
		rigidBodyQuery:       query.NewQuery(filter.Contains(component.RigidBody)),
		backgroundQuery:      query.NewQuery(filter.Contains(entity.BackgroundTag)),
		worldViewPortQuery:   query.NewQuery(filter.Contains(entity.WorldViewPortTag)),
		playerSlot:           query.NewQuery(filter.Contains(entity.SlotTag)),
		jellyBeanQuery:       query.NewQuery(filter.Contains(tag.JellyBean)),
		damageLabelQuery:     query.NewQuery(filter.Contains(entity.DamageLabelTag)),
		positionQuery:        query.NewQuery(filter.Contains(component.Position)),
		tilemap_cache:        nil,
	}
}

func (r *Render) Update(w donburi.World) {
	r.count++
	r.updateAnimatableActor(w)
	// r.updateDamageLabel(w)
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.tileMap(w, screen)
	r.debugRigidBodies(w, screen)
	r.animatableActor(w, screen)
	r.playerSlots(w, screen)
}

func (r *Render) updateAnimatableActor(w donburi.World) {
	r.animatableActorQuery.EachEntity(w, func(entry *donburi.Entry) {
		animation := component.GetAnimation(entry)
		anim := animation.GetCurrent(entry)
		if anim == nil {
			return
		}
		anim.NextFrame()
	})
}

// func (r Render) updateDamageLabel(w donburi.World) {
// 	r.damageLabelQuery.EachEntity(w, func(e *donburi.Entry) {
// 		t := component.GetTick(e)

// 		t.Value += 1
// 		if t.Value > t.EOL {
// 			w.Remove(e.Entity())
// 		}
// 	})
// }

func (r Render) renderDamageLabels(w donburi.World, screen *ebiten.Image) {
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.damageLabelQuery.EachEntity(w, func(e *donburi.Entry) {
		labelData := component.GetScreenLabel(e)
		position := component.GetPosition(e)
		txt := labelData.Label
		ebitenutil.DebugPrintAt(screen, txt, int(position.X)-int(worldViewLocationPos.X), int(position.Y)-int(worldViewLocationPos.Y))
	})
}

func (r Render) animatableActor(w donburi.World, screen *ebiten.Image) {
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.animatableActorQuery.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		animation := component.GetAnimation(entry)
		state := component.GetActorState(entry)

		anim := animation.GetCurrent(entry)
		if anim == nil {
			logrus.Error("unabled to get animation: ", state.GetCurrent())
			return
		}

		next := anim.GetFrame()
		if next == nil {
			logrus.Error("unabled to get frame")
			return
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X-position.CX-worldViewLocationPos.X, position.Y-position.CY-worldViewLocationPos.Y)
		screen.DrawImage(next, op)
	})
}

func (r *Render) debugRigidBodies(w donburi.World, screen *ebiten.Image) {
	if !config.Get().DebugEnabled {
		return
	}
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.rigidBodyQuery.EachEntity(w, func(entry *donburi.Entry) {
		p := component.GetPosition(entry)
		rb := component.GetRigidBody(entry)
		ebitenutil.DrawRect(screen, p.X-rb.L-worldViewLocationPos.X, p.Y-rb.T-worldViewLocationPos.Y, rb.GetWidth(), rb.GetHeight(), colornames.Red100)
	})
}

func (r *Render) tileMap(w donburi.World, screen *ebiten.Image) {
	worldViewLocation, _ := r.worldViewPortQuery.FirstEntity(w)
	worldViewLocationPos := component.GetPosition(worldViewLocation)

	r.backgroundQuery.EachEntity(w, func(entry *donburi.Entry) {

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0-worldViewLocationPos.X, 0-worldViewLocationPos.Y)

		if r.tilemap_cache == nil {
			tiles := component.GetTileMap(entry)

			if tiles.Map == nil {
				return
			}

			renderer, err := render.NewRenderer(tiles.Map)
			if err != nil {
				fmt.Printf("map unsupported for rendering: %s", err.Error())
				return
			}
			if err = renderer.RenderVisibleLayers(); err != nil {
				logrus.Error(err)
				return
			}
			r.tilemap_cache = ebiten.NewImageFromImage(renderer.Result)
		}

		if r.tilemap_cache == nil {
			return
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
