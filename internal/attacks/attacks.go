package attacks

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/pkg/gamemath"

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
					distance := gamemath.GetDistance([]float64{position.X, position.Y}, []float64{enemyPosition.X, enemyPosition.Y})
					if distance < closestDistance && w.Valid(pentry.Entity()) {
						closestDistance = distance
						closestEntry = pentry
					}
				}
			})

			if closestEntry != nil {
				e := gamemath.GetVector(position.X, position.Y)
				m := gamemath.GetVector(component.GetPosition(closestEntry).X, component.GetPosition(closestEntry).Y)
				entity.MakeRocketProjectile(w, position.X, position.Y, gamemath.GetHeading(e, m), am)
			}
		})
	}
}
