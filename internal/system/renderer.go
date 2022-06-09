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
type ViewPort interface {
	ViewPort() archetype.ViewPortTag
	GetPosition() *component.PositionData
}

type Renderer struct {
	World ecs.World
}

func NewRenderer(w ecs.World) Renderer {
	return Renderer{
		World: w,
	}
}

func (r Renderer) Update() {
	for _, entity := range ecs.FilterBy[Animatable](r.World) {
		r.updateAnimatable(entity)
	}
}

func (r Renderer) updateAnimatable(entity Animatable) {
	entity.IterFrame()
}

func (r Renderer) Draw(screen *ebiten.Image) {
	for _, entity := range ecs.FilterBy[Animatable](r.World) {
		r.animatable(screen, entity)
	}
}

func (r Renderer) animatable(screen *ebiten.Image, entity Animatable) {
	position := entity.GetPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.getWorldCoord(position))
	screen.DrawImage(entity.GetFrame(), op)

	r.healthBar(screen, entity)
}

func (r Renderer) healthBar(screen *ebiten.Image, entity Animatable) {
	if !ecs.IsType[Player](entity) {
		return
	}
	position := entity.GetPosition()
	health := entity.GetHealth()
	x, y := r.getWorldCoord(position)

	var marginBottom float64 = 35

	ebitenutil.DrawRect(screen, x, marginBottom+y, health.MAXHP/3, 3, colornames.Grey100)
	ebitenutil.DrawRect(screen, x, marginBottom+y, health.HP/3, 3, colornames.Red600)
}

func (r Renderer) getWorldCoord(position *component.PositionData) (float64, float64) {
	viewPort, err := ecs.FirstEntity[ViewPort](r.World)
	if err != nil {
		logrus.Errorf("viewport: %s", err)
		return position.X, position.Y
	}
	worldViewLocationPos := viewPort.GetPosition()
	return position.X - position.CX - worldViewLocationPos.X, position.Y - position.CY - worldViewLocationPos.Y
}
