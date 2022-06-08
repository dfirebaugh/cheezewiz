package component

import (
	"cheezewiz/pkg/animation"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
	Animations     map[string]*animation.Animation
	PrevUpdateTime time.Time
}

var Animation = donburi.NewComponentType(AnimationData{
	Animations: map[string]*animation.Animation{},
})

func GetAnimation(entry *donburi.Entry) *AnimationData {
	return (*AnimationData)(entry.Component(Animation))
}

func (a AnimationData) Get(label string) *animation.Animation {
	if _, ok := a.Animations[label]; !ok {
		logrus.Errorf("could not find an animation for this state: %d \n %#v \n", label, a.Animations)
		return a.Animations[string(DebugState)]
	}
	return a.Animations[label]
}

func (a AnimationData) GetCurrent(entry *donburi.Entry) *animation.Animation {
	state := GetActorState(entry)
	return a.Animations[string(state.GetCurrent())]
}
