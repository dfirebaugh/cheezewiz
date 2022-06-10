package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/pkg/attackgroup"
	"cheezewiz/pkg/ecs"

	"github.com/sirupsen/logrus"
)

type DamageBufferGroup struct {
	World ecs.World
}

func (d DamageBufferGroup) Update() {
	attackgroup.ApplyPlayerDamage(d)
	attackgroup.ApplyEnemyDamage(d)
}

func (d DamageBufferGroup) Apply(actor interface{}, amount float64) {
	var a archetype.Actor
	var ok bool
	if a, ok = actor.(archetype.Actor); !ok {
		return
	}

	health := a.GetHealth()
	state := a.GetState()
	logrus.Info("health: ", health.Current, " Origin Health ")

	if health.Current <= 0 {
		state.Set(component.DeathState)
		if !ecs.Is[archetype.Player](a) {
			d.World.Remove(a)
		}
		return
	}
	if health.Current > 0 {
		health.Current -= amount
		state.Set(component.HurtState)
	}
}
