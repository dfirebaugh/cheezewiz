package mediator

import (
	"cheezewiz/internal/system"

	"github.com/yohamta/donburi"
)

type Attack struct {
	C *system.Collision
	D *system.DamageBufferGroup
}

func (a Attack) AddPlayerDamage(destination *donburi.Entry, amount float64, origin *donburi.Entry) {
	a.D.AddPlayerDamage(destination, amount, origin)

}
func (a Attack) AddEnemyDamage(destination *donburi.Entry, amount float64, origin *donburi.Entry) {
	a.D.AddEnemyDamage(destination, amount, origin)

}
