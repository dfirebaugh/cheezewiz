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

type damageBufferGroup struct {
	PlayerDamage []damageToken
	EnemyDamage  []damageToken
}

func (d damageBufferGroup) AddPlayerDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry) {
	damage := damageToken{
		origin:      origin,
		amount:      amount,
		destination: reciever,
	}
	d.PlayerDamage = append(d.PlayerDamage, damage)
}

func (d damageBufferGroup) AddEnemyDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry) {
	damage := damageToken{
		origin:      origin,
		amount:      amount,
		destination: reciever,
	}
	d.EnemyDamage = append(d.EnemyDamage, damage)
}

func (d damageBufferGroup) ConsumeDamage() {
	for _, elem := range d.PlayerDamage { //Typecheck and warn!
		hc := component.GetHealth(elem.destination)
		hc.HP -= elem.amount

		if hc.HP <= 0 {
		}

		logrus.Infof("Death for entity %d", elem.destination.Id())

	}
	for _, elem := range d.EnemyDamage {
		hc := component.GetHealth(elem.destination)
		hc.HP -= elem.amount

		if hc.HP <= 0 {
		}

		logrus.Infof("Death for entity %d", elem.destination.Id())
	}

}

func (d damageBufferGroup) Update(w donburi.World) {
	d.ConsumeDamage()
}

func NewDamagebufferGroup() damageBufferGroup {
	buffergroup := damageBufferGroup{
		PlayerDamage: []damageToken{},
		EnemyDamage:  []damageToken{},
	}

	return buffergroup
}
