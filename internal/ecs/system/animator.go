package system

import (
	"cheezewiz/internal/ecs/component"
	"time"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type animator struct {
	Position   *component.Position
	Animation  *component.Animation
	Controller *component.Controller
}

func NewAnimator() gohan.System {
	return &animator{}
}

func (a *animator) Update(_ gohan.Entity) error {
	now := time.Now()

	if a.Controller.Controller.IsDownPressed() || a.Controller.Controller.IsDownPressed() || a.Controller.Controller.IsDownPressed() || a.Controller.Controller.IsDownPressed() {
		a.Animation.Walk.Update(now.Sub(a.Animation.PrevUpdateTime))
	} else {
		a.Animation.Still.Update(now.Sub(a.Animation.PrevUpdateTime))
	}
	a.Animation.PrevUpdateTime = now
	return nil
}

func (a *animator) Draw(entity gohan.Entity, screen *ebiten.Image) error {
	a.Animation.DrawOptions.Reset()

	a.Animation.DrawOptions.SetPos(a.Position.X-a.Animation.SpriteSize, a.Position.Y-a.Animation.SpriteSize)

	if a.Controller.Controller.IsDownPressed() || a.Controller.Controller.IsDownPressed() || a.Controller.Controller.IsDownPressed() || a.Controller.Controller.IsDownPressed() {
		a.Animation.Walk.Draw(screen, a.Animation.DrawOptions)
	} else {
		a.Animation.Still.Draw(screen, a.Animation.DrawOptions)
	}
	return nil
}
