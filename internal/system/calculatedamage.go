package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/attackgroup"
	"cheezewiz/pkg/ecs"

	"github.com/sirupsen/logrus"
)

type Actor interface {
	GetHealth() *component.Health
	GetActorState() *component.ActorState
}
type DamageBufferGroup struct {
	World ecs.World
}

func (d DamageBufferGroup) Update() {
	attackgroup.ApplyPlayerDamage(d)
	attackgroup.ApplyEnemyDamage(d)
}

func (DamageBufferGroup) Apply(actor interface{}, amount float64) {
	var a Actor
	var ok bool
	if a, ok = actor.(Actor); !ok {
		return
	}

	health := a.GetHealth()
	state := a.GetActorState()
	logrus.Info("health: ", health.Current, " Origin Health ")

	if health.Current == 0 {
		state.Set(component.DeathState)
		return
	}
	if health.Current > 0 {
		health.Current -= amount
		state.Set(component.HurtState)
	}

	// logrus.Infof("Death for entity %d", entry.Id())
}
