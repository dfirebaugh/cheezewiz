package component

import (
	"time"

	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type Action struct {
	Sprite         *ganim8.Sprite
	Animation      *ganim8.Animation
	PrevUpdateTime time.Time
}

type AnimationData struct {
	Walk  Action
	Idle  Action
	Hurt  Action
	Death Action
}

var Animation = donburi.NewComponentType(AnimationData{})

func GetAnimation(entry *donburi.Entry) *AnimationData {
	return (*AnimationData)(entry.Component(Animation))
}
