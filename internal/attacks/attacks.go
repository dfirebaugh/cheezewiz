package attacks

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/dentity"
	"cheezewiz/internal/tag"
	"cheezewiz/pkg/gamemath"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type attackGroup interface {
	AddEnemyDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry)
}

var CheeseMissile = func(world donburi.World) func() {
	w := world
	q := query.NewQuery(filter.Contains(tag.Player))
	destinationQuery := query.NewQuery(filter.Contains(tag.Enemy))
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

			if closestEntry == nil {
				return
			}

			launchProjectile(w, *position, *component.GetPosition(closestEntry))
		})
	}
}

func launchProjectile(w donburi.World, from component.PositionData, to component.PositionData) {
	e := gamemath.GetVector(from.X, from.Y)
	m := gamemath.GetVector(to.X, to.Y)
	dentity.MakeWithDirection(w, "./config/entities/rocket.entity.json", from.X, from.Y, gamemath.GetHeading(e, m))
}
