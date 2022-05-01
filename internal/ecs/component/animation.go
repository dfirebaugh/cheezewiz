package component

import (
	"time"

	"github.com/yohamta/ganim8/v2"
)

type Animation struct {
	Walk           *ganim8.Animation
	Still          *ganim8.Animation
	StillSprite    *ganim8.Sprite
	WalkSprite     *ganim8.Sprite
	SpriteSize     float64
	Grid           *ganim8.Grid
	PrevUpdateTime time.Time
	DrawOptions    *ganim8.DrawOptions
}
