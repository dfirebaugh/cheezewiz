package component

import (
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
	Available map[ActorStateType]ActorStateType
}

func (p *ActorState) Reset() {
	p.current = IdleState
}
func (as *ActorState) SetAvailable(animations map[ActorStateType]ActorStateType) {
	as.Available = map[ActorStateType]ActorStateType{}
	for label := range animations {
		as.Available[label] = label
	}
}
func (as *ActorState) Set(newState ActorStateType) {
	as.current = newState
}

func (as *ActorState) GetCurrent() ActorStateType {
	var current ActorStateType
	var ok bool

	if current, ok = as.Available[as.current]; !ok {
		logrus.Warnf("not able to lookup this actor's current state, falling back to debug state: %s", current)
		return DebugState
	}

	return ActorStateType(current)
}
