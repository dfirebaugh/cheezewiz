package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Animatable interface {
	GetFrame() *ebiten.Image
	GetPosition() *component.PositionData
	IterFrame()
	GetHealth() *component.HealthAspect
}

type Renderer struct {
	World            ecs.World
	animatableFilter ecs.FilterID
}

func NewRenderer(w ecs.World) Renderer {
	af := func(entity ecs.Entity) bool {
		if _, ok := entity.(Animatable); ok {
			return true
		}
		return false
	}
	return Renderer{
		World:            w,
		animatableFilter: w.MakeFilter(af),
	}
}

func (r Renderer) Update() {
	for _, entity := range r.World.FilterBy(r.animatableFilter) {
		r.updateAnimatable(entity)
	}
}

func (r Renderer) updateAnimatable(entity ecs.Entity) {
	var e Animatable
	var ok bool

	if e, ok = entity.(Animatable); !ok {
		println("entity doens't match update contract")
		return
	}

	e.IterFrame()
}

func (r Renderer) Draw(screen *ebiten.Image) {
	for _, entity := range r.World.FilterBy(r.animatableFilter) {
		r.animatable(screen, entity)
	}
}

func (r Renderer) animatable(screen *ebiten.Image, entity ecs.Entity) {
	var e Animatable
	var ok bool

	if e, ok = entity.(Animatable); !ok {
		println("entity doens't match draw contract")
		return
	}

	position := e.GetPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X, position.Y)
	screen.DrawImage(e.GetFrame(), op)

	r.healthBar(screen, entity)
}

func (r Renderer) healthBar(screen *ebiten.Image, entity ecs.Entity) {
	var e archetype.Actor
	var ok bool

	if e, ok = entity.(archetype.Actor); !ok {
		logrus.Info("doesn't match contract")
		return
	}

	position := e.GetPosition()
	health := e.GetHealth()
	// x, y := r.getWorldCoord(w, position)
	x, y := position.X, position.Y

	var marginBottom float64 = 35

	ebitenutil.DrawRect(screen, x, marginBottom+y, health.MAXHP/3, 3, colornames.Grey100)
	ebitenutil.DrawRect(screen, x, marginBottom+y, health.HP/3, 3, colornames.Red600)
}
