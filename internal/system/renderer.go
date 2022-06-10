package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/cache"
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/colornames"
)

type Renderer struct {
	World ecs.World
}

func NewRenderer(w ecs.World) Renderer {
	return Renderer{
		World: w,
	}
}

func (r *Renderer) Update() {
	var animatables []archetype.Animatable
	var ok bool

	if animatables, ok = cache.GetAnimatables(r.World); !ok {
		return
	}
	for _, entity := range animatables {
		r.updateAnimatable(entity)
	}
}

func (r Renderer) updateAnimatable(entity archetype.Animatable) {
	entity.IterFrame()
}

func (r Renderer) Draw(screen *ebiten.Image) {
	var animatables []archetype.Animatable
	var ok bool

	r.debug(screen)

	if animatables, ok = cache.GetAnimatables(r.World); !ok {
		return
	}
	for _, entity := range animatables {
		r.animatable(screen, entity)
	}
}

func (r Renderer) animatable(screen *ebiten.Image, entity archetype.Animatable) {
	position := entity.GetPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.getWorldCoord(position))
	screen.DrawImage(entity.GetFrame(), op)

	r.healthBar(screen, entity)
}

func (r Renderer) healthBar(screen *ebiten.Image, entity archetype.Animatable) {
	if !ecs.Is[archetype.Player](entity) {
		return
	}
	position := entity.GetPosition()
	health := entity.GetHealth()
	x, y := r.getWorldCoord(position)

	var marginBottom float64 = 35

	ebitenutil.DrawRect(screen, x, marginBottom+y, health.Max/3, 3, colornames.Grey)
	ebitenutil.DrawRect(screen, x, marginBottom+y, health.Current/3, 3, colornames.Red)
}

func (r Renderer) getWorldCoord(position *component.Position) (float64, float64) {
	viewPort, err := ecs.FirstEntity[archetype.ViewPort](r.World)
	if err != nil {
		logrus.Errorf("viewport: %s", err)
		return position.X, position.Y
	}
	worldViewLocationPos := viewPort.GetPosition()
	return position.X - position.CX - worldViewLocationPos.X, position.Y - position.CY - worldViewLocationPos.Y
}

func (r Renderer) debug(screen *ebiten.Image) {
	if !config.Get().DebugEnabled {
		return
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f -- entities: %d\n", ebiten.CurrentFPS(), r.World.Count()))

	var collidables []archetype.Collidable
	var ok bool
	if collidables, ok = cache.GetCollidables(r.World); !ok {
		return
	}
	for _, c := range collidables {
		p := c.GetPosition()
		rb := c.GetRigidBody()
		x, y := r.getWorldCoord(p)
		ebitenutil.DrawRect(screen, x, y, rb.GetWidth(), rb.GetHeight(), colornames.Aliceblue)
	}
}
