package system

import (
	"cheezewiz/internal/component"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type damageToken struct {
	origin      *donburi.Entry
	amount      float64
	destination *donburi.Entry
}

type DamageBufferGroup struct {
	PlayerDamage []damageToken
	EnemyDamage  []damageToken
}

func (d *DamageBufferGroup) AddPlayerDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry) {
	damage := damageToken{
		origin:      origin,
		amount:      amount,
		destination: reciever,
	}
	d.PlayerDamage = append(d.PlayerDamage, damage)

}

func (d *DamageBufferGroup) AddEnemyDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry) {
	damage := damageToken{
		origin:      origin,
		amount:      amount,
		destination: reciever,
	}
	d.EnemyDamage = append(d.EnemyDamage, damage)
}

func (d *DamageBufferGroup) ConsumeDamage() {
	for _, elem := range d.PlayerDamage { //Typecheck and warn!
		playerHeath := component.GetHealth(elem.destination)
		state := component.GetPlayerState(elem.destination)
		logrus.Info("players' health: ", playerHeath.HP, " Origin Health ")

		if playerHeath.HP > 0 {
			playerHeath.HP -= elem.amount
		}

		state.Current = component.HurtState

		logrus.Infof("Death for entity %d", elem.destination.Id())
	}

	d.PlayerDamage = nil

	for _, elem := range d.EnemyDamage {
		hc := component.GetHealth(elem.destination)
		hc.HP -= elem.amount

		logrus.Infof("Death for entity %d", elem.destination.Id())
	}

}

func (d *DamageBufferGroup) Update(w donburi.World) {
	d.ConsumeDamage()
}

func NewDamagebufferGroup() DamageBufferGroup {
	buffergroup := DamageBufferGroup{
		PlayerDamage: []damageToken{},
		EnemyDamage:  []damageToken{},
	}

	return buffergroup
}
