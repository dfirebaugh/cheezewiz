package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/atedja/go-vector"
)

type EnemyControl struct {
	query            *query.Query
	destinationQuery *query.Query
}

const enemySpeed = 0.5

func NewEnemyControl() *EnemyControl {
	return &EnemyControl{
		query:            query.NewQuery(filter.Contains(entity.EnemyTag)),
		destinationQuery: query.NewQuery(filter.Contains(entity.PlayerTag)),
	}
}

func (e EnemyControl) Update(w donburi.World) {
	e.query.EachEntity(w, func(entry *donburi.Entry) {
		entityPosition := component.GetPosition(entry)

		var closestEntry *donburi.Entry
		var closestDistance float64 = 100000000

		e.destinationQuery.EachEntity(w, func(pentry *donburi.Entry) {
			playerPosition := component.GetPosition(pentry)
			if closestEntry != nil {
				closestEntry = pentry
			} else {
				distance := math.Sqrt(math.Pow(entityPosition.X-playerPosition.X, 2) + math.Pow(entityPosition.Y-playerPosition.Y, 2))
				if distance < closestDistance {
					closestDistance = distance
					closestEntry = pentry
				}
			}

		})

		if closestEntry != nil {

			p := vector.NewWithValues([]float64{component.GetPosition(closestEntry).X, component.GetPosition(closestEntry).Y})

			e := vector.NewWithValues([]float64{entityPosition.X, entityPosition.Y})

			r := vector.Unit(vector.Subtract(p, e))

			r.Scale(enemySpeed)

			newloc := vector.Add(e, r)

			entityPosition.X = newloc[0]
			entityPosition.Y = newloc[1]

		}

	})

}
