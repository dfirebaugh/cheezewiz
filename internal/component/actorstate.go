package component

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
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

type ActorStateData struct {
	current   ActorStateType
	Available map[string]string
	stateMut  sync.RWMutex
}

var ActorState = donburi.NewComponentType(ActorStateData{current: IdleState})

func GetActorState(entry *donburi.Entry) *ActorStateData {
	return (*ActorStateData)(entry.Component(ActorState))
}

func (p *ActorStateData) Reset() {
	p.current = IdleState
}
func (as *ActorStateData) SetAvailable(animations map[string]string) {
	as.stateMut.Lock()
	defer as.stateMut.Unlock()

	as.Available = map[string]string{}
	for label := range animations {
		as.Available[label] = label
	}
}
func (as *ActorStateData) Set(newState interface{}) {
	as.stateMut.Lock()
	defer as.stateMut.Unlock()

	var ok bool
	as.current, ok = newState.(ActorStateType)

	if !ok {
		logrus.Warn("this state is not defined: ", newState)
		as.Set(DebugState)
	}
}

func (as *ActorStateData) GetCurrent() ActorStateType {
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
