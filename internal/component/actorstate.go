package component

import (
	"github.com/sirupsen/logrus"
)

type StateType string

const (
	DebugState     StateType = "debug"
	IdleState      StateType = "idle"
	WalkingState   StateType = "walk"
	AttackingState StateType = "attack"
	HurtState      StateType = "hurt"
	DeathState     StateType = "death"
)

type State struct {
	current   StateType
	Available map[StateType]StateType
}

func (p *State) Reset() {
	p.current = IdleState
}
func (as *State) SetAvailable(animations map[StateType]StateType) {
	as.Available = map[StateType]StateType{}
	for label := range animations {
		as.Available[label] = label
	}
}
func (as *State) Set(newState StateType) {
	as.current = newState
}

func (as *State) GetCurrent() StateType {
	var current StateType
	var ok bool

	if current, ok = as.Available[as.current]; !ok {
		logrus.Warnf("not able to lookup this actor's current state, falling back to debug state: %s", current)
		return DebugState
	}

	return StateType(current)
}
