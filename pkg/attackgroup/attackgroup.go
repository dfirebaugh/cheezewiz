package attackgroup

type Actor interface {
}

type DamageApplier interface {
	Apply(actor interface{}, amount float64)
}

type DamageToken struct {
	Origin      Actor
	Amount      float64
	Destination Actor
}

type DamageBuffergroup struct {
	PlayerDamage []DamageToken
	EnemyDamage  []DamageToken
}

var Buffergroup = DamageBuffergroup{
	PlayerDamage: []DamageToken{},
	EnemyDamage:  []DamageToken{},
}

func AddPlayerDamage(reciever Actor, amount float64, origin Actor) {
	damage := DamageToken{
		Origin:      origin,
		Amount:      amount,
		Destination: reciever,
	}
	Buffergroup.PlayerDamage = append(Buffergroup.PlayerDamage, damage)
}

func AddEnemyDamage(reciever Actor, amount float64, Origin Actor) {
	damage := DamageToken{
		Origin:      Origin,
		Amount:      amount,
		Destination: reciever,
	}
	Buffergroup.EnemyDamage = append(Buffergroup.EnemyDamage, damage)
}

func ApplyPlayerDamage(dmg DamageApplier) {
	for _, elem := range Buffergroup.PlayerDamage {
		dmg.Apply(elem.Destination, elem.Amount)
	}
	// clear our the PlayerDamage buffer
	Buffergroup.PlayerDamage = []DamageToken{}
}

func ApplyEnemyDamage(dmg DamageApplier) {
	for _, elem := range Buffergroup.EnemyDamage {
		dmg.Apply(elem.Destination, elem.Amount)
	}
	// clear our the EnemyDamage buffer
	Buffergroup.EnemyDamage = []DamageToken{}
}
