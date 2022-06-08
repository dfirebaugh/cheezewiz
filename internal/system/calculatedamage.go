package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/attackgroup"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type DamageBufferGroup struct{}

func (d DamageBufferGroup) Update(w donburi.World) {
	attackgroup.ApplyPlayerDamage(d)
	attackgroup.ApplyEnemyDamage(d)
}

func (DamageBufferGroup) Apply(w donburi.World, entry *donburi.Entry, amount float64) {
	health := component.GetHealth(entry)
	state := component.GetActorState(entry)
	logrus.Info("health: ", health.HP, " Origin Health ")

	if health.HP > 0 {
		health.HP -= amount
	}

	state.Set(component.HurtState)
	logrus.Infof("Death for entity %d", entry.Id())
}
