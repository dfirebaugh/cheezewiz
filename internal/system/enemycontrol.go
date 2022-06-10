package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/gamemath"
	"math"

	"github.com/atedja/go-vector"
)

type Enemy interface {
	GetPosition() *component.Position
	GetHealth() *component.Health
	GetEnemyTag() ecs.Tag
}

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
	for _, ent := range ecs.FilterBy[Enemy](e.world) {
		isDead := e.updateHealth(ent)
		if isDead {
			return
		}

		e.updatePosition(ent)
	}
}

// updateHealh returns true if the enemy has died
func (e EnemyControl) updateHealth(enemy Enemy) bool {
	health := enemy.GetHealth()
	position := enemy.GetPosition()
	if health.Current <= 0 {
		// w.Remove(entry.Entity())
		entity.MakeRandEntity(e.world, []string{
			"entities/jellybeangreen.entity.json",
			"entities/jellybeanpink.entity.json",
			"entities/jellybeanblue.entity.json",
			"entities/jellybeanrainbow.entity.json",
		}, position.X-position.CX, position.Y-position.CY)
		return true
	}
	return false
}

func (e EnemyControl) updatePosition(enemy Enemy) {
	if enemy == nil {
		return
	}
	position := enemy.GetPosition()
	var closestPlayer Player
	var closestDistance float64 = 100000000

	for _, player := range ecs.FilterBy[Player](e.world) {
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

func (ec EnemyControl) moveTowardPlayer(player Player, enemy Enemy) {
	if player == nil || enemy == nil {
		return
	}
	position := enemy.GetPosition()
	playerPosition := player.GetPosition()

	p := vector.NewWithValues([]float64{playerPosition.X, playerPosition.Y})

	e := vector.NewWithValues([]float64{position.X, position.Y})

	r := vector.Unit(vector.Subtract(p, e))

	speed := enemySpeed
	if ec.isOverPositionLimit(e, p, 100) {
		// as they get closer, slow down a bit
		speed = .2
	}
	if ec.isOverPositionLimit(e, p, 23) {
		return
	}
	r.Scale(speed)

	newloc := vector.Add(e, r)

	position.Update(newloc[0], newloc[1])
}

func (ec EnemyControl) isOverPositionLimit(a []float64, b []float64, limit float64) bool {
	return gamemath.GetDistance(a, b) < limit
}
