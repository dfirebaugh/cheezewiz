package component

import "github.com/yohamta/donburi"

type ActorStateType uint

const (
	IdleState ActorStateType = iota
	WalkingState
	AttackingState
	HurtState
	DeathState
)

type ActorStateData struct {
	Current ActorStateType
}

var ActorState = donburi.NewComponentType(ActorStateData{Current: IdleState})

func GetActorState(entry *donburi.Entry) *ActorStateData {
	return (*ActorStateData)(entry.Component(ActorState))
}

func (p *ActorStateData) Reset() {
	p.Current = IdleState
}

func (p *ActorStateData) Set(newState ActorStateType) {
	p.Current = newState
}
