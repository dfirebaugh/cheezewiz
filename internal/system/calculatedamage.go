package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/attackgroup"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type DamageBufferGroup struct{}

func (d DamageBufferGroup) Update(w donburi.World) {
	for _, elem := range attackgroup.Buffergroup.PlayerDamage { //Typecheck and warn!
		playerHeath := component.GetHealth(elem.Destination)
		state := component.GetActorState(elem.Destination)
		logrus.Info("players' health: ", playerHeath.HP, " Origin Health ")

		if playerHeath.HP > 0 {
			playerHeath.HP -= elem.Amount
		}

		state.Set(component.HurtState)

		logrus.Infof("Death for entity %d", elem.Destination.Id())
	}

	// clear our the PlayerDamage buffer
	attackgroup.Buffergroup.PlayerDamage = []attackgroup.DamageToken{}

	for _, elem := range attackgroup.Buffergroup.EnemyDamage {
		attackgroup.ApplyDamageToEnemy(attackgroup.Buffergroup.World, elem.Destination, elem.Amount)
	}

	// clear our the EnemyDamage buffer
	attackgroup.Buffergroup.EnemyDamage = []attackgroup.DamageToken{}
}
