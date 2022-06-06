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
	world        donburi.World
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
		state := component.GetActorState(elem.destination)
		logrus.Info("players' health: ", playerHeath.HP, " Origin Health ")

		if playerHeath.HP > 0 {
			playerHeath.HP -= elem.amount
		}

		state.Current = component.HurtState

		logrus.Infof("Death for entity %d", elem.destination.Id())
	}

	// clear our the PlayerDamage buffer
	d.PlayerDamage = []damageToken{}

	for _, elem := range d.EnemyDamage {
		applyDamageToEnemy(d.world, elem.destination, elem.amount)
	}

	// clear our the EnemyDamage buffer
	d.EnemyDamage = []damageToken{}
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

func applyDamageToEnemy(w donburi.World, enemy *donburi.Entry, amount float64) {
	if enemy == nil {
		println("enemy is not valid")
		return
	}

	hc := component.GetHealth(enemy)
	hc.HP -= amount

	// logrus.Infof("Damage delt for entity %#v with %d dmg", enemy.Id(), amount)
	// remove enemy from damage buffer
	enemy = nil
}
