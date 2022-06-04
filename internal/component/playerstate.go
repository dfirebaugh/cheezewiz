package component

import "github.com/yohamta/donburi"

type playerstate uint

const (
	IdleState playerstate = iota
	WalkingState
	AttackingState
	HurtState
	DeathState
)

type PlayerState struct {
	Current playerstate
}

var State = donburi.NewComponentType(PlayerState{Current: IdleState})

func GetPlayerState(entry *donburi.Entry) *PlayerState {
	return (*PlayerState)(entry.Component(State))
}

func (p *PlayerState) ResetState() {
	p.Current = IdleState
}
