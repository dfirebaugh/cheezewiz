package component

import (
	"cheezewiz/pkg/animation"
	"time"
)

type Animation struct {
	Animation      map[StateType]*animation.Animation
	PrevUpdateTime time.Time
}
