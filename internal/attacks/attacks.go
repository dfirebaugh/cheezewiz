package attacks

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/gamemath"

	"github.com/atedja/go-vector"
)

type attackGroup interface {
	AddEnemyDamage(reciever archetype.Actor, amount float64, origin archetype.Actor)
}

type Directionable interface {
	GetDirection() *component.Direction
}

var CheeseMissile = func(world ecs.World) func() {
	w := world
	return func() {
		for handle, player := range ecs.FilterBy[archetype.Player](w) {
			findHeading(w, player, handle)
		}
	}
}

func findHeading(w ecs.World, player archetype.Player, playerHandle int) {
	position := player.GetPosition()
	state := player.GetActorState()

	enemies := map[int]vector.Vector{}

	for handle, actor := range ecs.FilterBy[archetype.Actor](w) {
		if handle == playerHandle {
			continue
		}
		p := actor.GetPosition()
		enemies[handle] = vector.NewWithValues([]float64{p.X, p.Y})
	}

	closestHandle := gamemath.GetClosest(vector.NewWithValues([]float64{position.X, position.Y}), enemies)
	if closestHandle == playerHandle {
		return
	}
	// closestEnemy, ok := w.EntityMap[closestHandle].(archetype.Actor)
	// if !ok {
	// 	return
	// }
	// launchProjectile(w, *position, *closestEnemy.GetPosition())
	state.Set(component.AttackingState)
}

func launchProjectile(w ecs.World, from component.Position, to component.Position) {
	e := gamemath.GetVector(from.X, from.Y)
	m := gamemath.GetVector(to.X, to.Y)
	entity.MakeWithDirection(w, "entities/rocket.entity.json", from.X, from.Y, gamemath.GetHeading(e, m))
}
