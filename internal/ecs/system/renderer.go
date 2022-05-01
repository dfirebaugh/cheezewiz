package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/ecs/component"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type renderer struct {
	Position *component.Position
	Asset    *component.Asset

	op *ebiten.DrawImageOptions
}

func NewRenderer() gohan.System {
	return &renderer{
		op: &ebiten.DrawImageOptions{},
	}
}

func (s *renderer) Update(_ gohan.Entity) error {
	return gohan.ErrUnregister
}

func (s *renderer) Draw(entity gohan.Entity, screen *ebiten.Image) error {
	c := config.Get()
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.Position.X-c.SpriteSize, s.Position.Y-c.SpriteSize)
	screen.DrawImage(s.Asset.Image, s.op)
	return nil
}
