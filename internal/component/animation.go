package component

import (
	"cheezewiz/assets"
	"time"

	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type Action struct {
	Sprite    *ganim8.Sprite
	Animation *ganim8.Animation
}

type AnimationData struct {
	Animations     map[ActorStateType]Action
	PrevUpdateTime time.Time
}

var Animation = donburi.NewComponentType(AnimationData{
	Animations: map[ActorStateType]Action{},
})

func GetAnimation(entry *donburi.Entry) *AnimationData {
	return (*AnimationData)(entry.Component(Animation))
}

func (a AnimationData) Get(state ActorStateType) Action {
	if _, ok := a.Animations[state]; !ok {

		return ToAction(assets.GetRaddishhWalking())
	}
	return a.Animations[state]
}

func ToAction(spr *ganim8.Sprite, a *ganim8.Animation) Action {
	return Action{
		Sprite:    spr,
		Animation: a,
	}
}
