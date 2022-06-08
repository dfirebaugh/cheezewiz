package attackgroup

import (
	"fmt"

	"github.com/yohamta/donburi"
)

type DamageToken struct {
	Origin      *donburi.Entry
	Amount      float64
	Destination *donburi.Entry
}

type DamageBuffergroup struct {
	PlayerDamage []DamageToken
	EnemyDamage  []DamageToken
	World        donburi.World
}

var Buffergroup = DamageBuffergroup{
	PlayerDamage: []DamageToken{},
	EnemyDamage:  []DamageToken{},
}

func AddPlayerDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry) {
	damage := DamageToken{
		Origin:      origin,
		Amount:      amount,
		Destination: reciever,
	}
	Buffergroup.PlayerDamage = append(Buffergroup.PlayerDamage, damage)
}

func AddEnemyDamage(reciever *donburi.Entry, amount float64, Origin *donburi.Entry) {
	damage := DamageToken{
		Origin:      Origin,
		Amount:      amount,
		Destination: reciever,
	}
	Buffergroup.EnemyDamage = append(Buffergroup.EnemyDamage, damage)
}

func ApplyDamageToEnemy(w donburi.World, enemy *donburi.Entry, amount float64) {
	if enemy == nil {
		println("enemy is not valid")
		return
	}

	fmt.Printf("apply %f dmg to enemy", amount)

	// hc := component.GetHealth(enemy)
	// hc.HP -= amount

	// remove enemy from damage buffer
	enemy = nil
}
