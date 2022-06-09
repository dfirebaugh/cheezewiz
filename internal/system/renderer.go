package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Animatable interface {
	GetFrame() *ebiten.Image
	GetPosition() *component.Position
	IterFrame()
	GetHealth() *component.Health
}
type ViewPort interface {
	ViewPort() ecs.Tag
	GetPosition() *component.Position
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
	r.collidable(screen)
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

	ebitenutil.DrawRect(screen, x, marginBottom+y, health.Max/3, 3, colornames.Grey100)
	ebitenutil.DrawRect(screen, x, marginBottom+y, health.Current/3, 3, colornames.Red600)
}

func (r Renderer) getWorldCoord(position *component.Position) (float64, float64) {
	viewPort, err := ecs.FirstEntity[ViewPort](r.World)
	if err != nil {
		logrus.Errorf("viewport: %s", err)
		return position.X, position.Y
	}
	worldViewLocationPos := viewPort.GetPosition()
	return position.X - position.CX - worldViewLocationPos.X, position.Y - position.CY - worldViewLocationPos.Y
}

func (r Renderer) collidable(screen *ebiten.Image) {
	if !config.Get().DebugEnabled {
		return
	}

	for _, c := range ecs.FilterBy[Collidable](r.World) {
		p := c.GetPosition()
		rb := c.GetRigidBody()
		x, y := r.getWorldCoord(p)
		ebitenutil.DrawRect(screen, x, y, rb.GetWidth(), rb.GetHeight(), colornames.Red100)
	}
}
