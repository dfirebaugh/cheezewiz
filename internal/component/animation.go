package component

import (
	"cheezewiz/pkg/animation"
	"time"

	"github.com/sirupsen/logrus"
)

type Animation struct {
	Animation      map[string]*animation.Animation
	PrevUpdateTime time.Time
}

func (a Animation) Get(label string) *animation.Animation {
	if _, ok := a.Animation[label]; !ok {
		logrus.Errorf("could not find an animation for this state: %d \n %#v \n", label, a.Animation)
		return a.Animation[string(DebugState)]
	}
	return a.Animation[label]
}

// func (a Animation) GetCurrent() *animation.Animation {
// 	// state := GetActorState(entry)
// 	// return a.Animations[string(state.GetCurrent())]
// }
