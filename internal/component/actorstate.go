package component

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type ActorStateType string

const (
	DebugState     ActorStateType = "debug"
	IdleState      ActorStateType = "idle"
	WalkingState   ActorStateType = "walk"
	AttackingState ActorStateType = "attack"
	HurtState      ActorStateType = "hurt"
	DeathState     ActorStateType = "death"
)

type ActorState struct {
	current   ActorStateType
	Available map[string]string
	stateMut  sync.RWMutex
}

func (p *ActorState) Reset() {
	p.current = IdleState
}
func (as *ActorState) SetAvailable(animations map[string]string) {
	as.stateMut.Lock()
	defer as.stateMut.Unlock()

	as.Available = map[string]string{}
	for label := range animations {
		as.Available[label] = label
	}
}
func (as *ActorState) Set(newState interface{}) {
	as.stateMut.Lock()
	defer as.stateMut.Unlock()

	var ok bool
	as.current, ok = newState.(ActorStateType)

	if !ok {
		logrus.Warn("this state is not defined: ", newState)
		as.Set(DebugState)
	}
}

func (as *ActorState) GetCurrent() ActorStateType {
	as.stateMut.Lock()
	defer as.stateMut.Unlock()

	var current string
	var ok bool

	if current, ok = as.Available[string(as.current)]; !ok {
		logrus.Warnf("not able to lookup this actor's current state, falling back to debug state: %s", current)
		return DebugState
	}

	return ActorStateType(current)
}
