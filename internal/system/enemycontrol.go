package system

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/cache"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/gamemath"
	"math"

	"github.com/atedja/go-vector"
)

type EnemyControl struct {
	world ecs.World
}

const enemySpeed = 0.5

func NewEnemyControl(w ecs.World) *EnemyControl {
	return &EnemyControl{
		world: w,
	}
}

func (e EnemyControl) Update() {
	var enemies []archetype.Enemy
	var ok bool
	if enemies, ok = cache.GetEnemies(e.world); !ok {
		return
	}
	for _, ent := range enemies {
		e.updatePosition(ent)
	}
}

func (e EnemyControl) updatePosition(enemy archetype.Enemy) {
	position := enemy.GetPosition()
	var closestPlayer archetype.Player
	var closestDistance float64 = 100000000

	if enemy.GetHealth().Current <= 0 {
		return
	}

	var players []archetype.Player
	var ok bool
	if players, ok = cache.GetPlayers(e.world); !ok {
		return
	}
	for _, player := range players {
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

	e.moveTowardPlayer(closestPlayer, enemy)
}

func (ec EnemyControl) moveTowardPlayer(player archetype.Player, enemy archetype.Enemy) {
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

func (ec EnemyControl) isOverPositionLimit(a []float64, b []float64, limit float64) bool {
	return gamemath.GetDistance(a, b) < limit
}
