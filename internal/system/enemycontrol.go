package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"

	"cheezewiz/pkg/gamemath"
	"math"

	"github.com/atedja/go-vector"
)

type EnemyController struct {
	// get references to players ahead of time
	//  for performance reasons.  We may need to update this list if a new player joins
	playerHandles []world.EntityHandle
}

const enemySpeed = 0.5

func (e *EnemyController) findPlayers() {
	e.playerHandles = query.Get(world.Instance, filter.GetPlayers)
}
func (e *EnemyController) Update() {
	if len(e.playerHandles) == 0 {
		e.findPlayers()
		return
	}

	query.Each(world.Instance, filter.GetEnemies, func(handle world.EntityHandle) {
		e.updatePosition(world.Instance.GetEntity(handle))
	})
}

func (e *EnemyController) updatePosition(enemy entity.Entity) {
	position := enemy.GetPosition()
	if enemy.GetHealth().Current <= 0 {
		return
	}

	e.moveTowardPlayer(e.findClosestPlayer(position), enemy)
}

func (e *EnemyController) findClosestPlayer(position *component.Position) entity.Entity {
	var closestPlayer entity.Entity
	var closestDistance float64 = 100000000
	for _, handle := range e.playerHandles {
		player := world.Instance.GetEntity(handle)
		playerPosition := player.GetPosition()
		if closestPlayer != nil {
			closestPlayer = player
		} else {
			distance := math.Sqrt(math.Pow(position.X-playerPosition.X, 2) + math.Pow(position.Y-playerPosition.Y, 2))
			if distance < closestDistance {
				closestDistance = distance
				closestPlayer = player
			}
		}
	}

	return closestPlayer
}

func (ec *EnemyController) moveTowardPlayer(player entity.Entity, enemy entity.Entity) {
	if player == nil || enemy == nil {
		return
	}
	position := enemy.GetPosition()
	playerPosition := player.GetPosition()

	p := vector.NewWithValues([]float64{playerPosition.X, playerPosition.Y})

	e := vector.NewWithValues([]float64{position.X, position.Y})

	r := vector.Unit(vector.Subtract(p, e))

	speed := enemySpeed
	if ec.isOverPositionLimit(e, p, 80) {
		// as they get closer, slow down a bit
		speed = .2
	}
	if ec.isOverPositionLimit(e, p, 28) {
		return
	}
	r.Scale(speed)

	newloc := vector.Add(e, r)

	position.Update(newloc[0], newloc[1])
}

func (ec *EnemyController) isOverPositionLimit(a []float64, b []float64, limit float64) bool {
	return gamemath.GetDistance(a, b) < limit
}
