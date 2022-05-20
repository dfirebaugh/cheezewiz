package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/ecs/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderable interface {
	GetAsset() *component.Asset
	GetPosition() *component.Position
	GetDrawOptions() *ebiten.DrawImageOptions
}

type Renderer struct {
	Level *component.Level
}

func (r Renderer) Render(screen *ebiten.Image) {
	c := config.Get()
	for _, id := range r.Level.Entities {
		if _, ok := r.Level.EntityMap[id].(Renderable); !ok {
			println("entity doens't match contract")
			continue
		}

		entity := r.Level.EntityMap[id].(Renderable)

		entity.GetDrawOptions().GeoM.Reset()
		entity.GetDrawOptions().GeoM.Translate(entity.GetPosition().X-c.SpriteSize, entity.GetPosition().Y-c.SpriteSize)
		screen.DrawImage(entity.GetAsset().Image, entity.GetDrawOptions())
	}
}
