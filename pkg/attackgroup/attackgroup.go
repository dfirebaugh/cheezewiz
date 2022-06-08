package attackgroup

import (
	"github.com/yohamta/donburi"
)

type DamageApplier interface {
	Apply(w donburi.World, entry *donburi.Entry, amount float64)
}

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

func ApplyPlayerDamage(dmg DamageApplier) {
	for _, elem := range Buffergroup.PlayerDamage {
		dmg.Apply(Buffergroup.World, elem.Destination, elem.Amount)
	}
	// clear our the PlayerDamage buffer
	Buffergroup.PlayerDamage = []DamageToken{}
}

func ApplyEnemyDamage(dmg DamageApplier) {
	for _, elem := range Buffergroup.EnemyDamage {
		dmg.Apply(Buffergroup.World, elem.Destination, elem.Amount)
	}
	// clear our the EnemyDamage buffer
	Buffergroup.EnemyDamage = []DamageToken{}
}
