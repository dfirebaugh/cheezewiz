package attacks

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/ecs/adapter"
	"cheezewiz/internal/entity"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/gamemath"

	"github.com/atedja/go-vector"
)

type Actor interface {
	GetPosition() *component.Position
}
type Player interface {
	GetPosition() *component.Position
	GetState() *component.State
	GetPlayerTag() ecs.Tag
}

type attackGroup interface {
	AddEnemyDamage(reciever Actor, amount float64, origin Actor)
}

type Directionable interface {
	GetDirection() *component.Direction
}

var CheeseMissile = func(world adapter.Adapter) func() {
	w := world
	return func() {
		for handle, player := range ecs.FilterMapBy[Player](w.GetWorld()) {
			findHeading(w, player, handle)
		}
	}
}

func findHeading(w adapter.Adapter, player Player, playerHandle ecs.EntityHandle) {
	position := player.GetPosition()
	state := player.GetState()
	if state.GetCurrent() == component.DeathState {
		return
	}

	enemies := map[ecs.EntityHandle]vector.Vector{}

	for handle, actor := range ecs.FilterMapBy[Actor](w.GetWorld()) {
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
	closestEnemy, ok := w.GetEntity(closestHandle).(Actor)
	if !ok {
		return
	}
	launchProjectile(w, *position, *closestEnemy.GetPosition())
	state.Set(component.AttackingState)
}

func launchProjectile(w adapter.Adapter, from component.Position, to component.Position) {
	e := gamemath.GetVector(from.X, from.Y)
	m := gamemath.GetVector(to.X, to.Y)
	entity.MakeWithDirection(w, "entities/rocket.entity.json", from.X, from.Y, gamemath.GetHeading(e, m))
}
