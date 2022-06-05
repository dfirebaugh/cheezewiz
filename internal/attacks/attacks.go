package attacks

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"math"

	"github.com/atedja/go-vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type attackGroup interface {
	AddEnemyDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry)
}

var CheeseMissile = func(world donburi.World, attackMediator attackGroup) func() {
	w := world
	am := attackMediator
	q := query.NewQuery(filter.Contains(entity.PlayerTag))
	destinationQuery := query.NewQuery(filter.Contains(entity.EnemyTag))
	return func() {
		q.EachEntity(w, func(e *donburi.Entry) {
			position := component.GetPosition(e)
			state := component.GetActorState(e)
			state.Set(component.AttackingState)
			var closestEntry *donburi.Entry
			var closestDistance float64 = 100000000

			destinationQuery.EachEntity(w, func(pentry *donburi.Entry) {
				enemyPosition := component.GetPosition(pentry)
				if closestEntry == nil && w.Valid(pentry.Entity()) {
					closestEntry = pentry
				} else {
					distance := math.Sqrt(math.Pow(position.X-enemyPosition.X, 2) + math.Pow(position.Y-enemyPosition.Y, 2))
					if distance < closestDistance && w.Valid(pentry.Entity()) {
						closestDistance = distance
						closestEntry = pentry
					}
				}

			})

			if closestEntry != nil {

				e := vector.NewWithValues([]float64{component.GetPosition(closestEntry).X, component.GetPosition(closestEntry).Y})

				m := vector.NewWithValues([]float64{position.X, position.Y})

				r := vector.Unit(vector.Subtract(m, e))

				entity.MakeRocketProjectile(w, position.X, position.Y, math.Atan2(r[1], r[0]), am)

			}

		})

	}
}
