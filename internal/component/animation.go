package component

import (
	"cheezewiz/pkg/animation"
	"time"
)

type Animation struct {
	Animation      map[ActorStateType]*animation.Animation
	PrevUpdateTime time.Time
}
