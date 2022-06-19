package attacks

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/tag"
	"cheezewiz/internal/world"
)

type attackGroup interface {
	AddEnemyDamage(reciever entity.Entity, amount float64, origin entity.Entity)
}

type Directionable interface {
	GetDirection() *component.Direction
}

var CheeseMissile = func(wl world.World) func() {
	w := wl
	return func() {
		w.EachEntity(func(handle world.EntityHandle) {
			if !w.GetEntity(handle).HasTag(tag.Player) {
				return
			}
			player := w.GetEntity(handle)

			findHeading(w, player, handle)
		})
	}
}

func findHeading(w world.World, player entity.Entity, playerHandle world.EntityHandle) {
	// position := player.GetPosition()
	// state := player.GetState()
	// if state.GetCurrent() == component.DeathState {
	// 	return
	// }

	// enemies := map[world.EntityHandle]vector.Vector{}

	// w.EachEntity(func(handle world.EntityHandle) {
	// 	if handle == playerHandle {
	// 		return
	// 	}
	// 	p := w.GetEntity(handle).GetPosition()
	// 	enemies[handle] = vector.NewWithValues([]float64{p.X, p.Y})
	// })

	// if len(enemies) == 0 {
	// 	return
	// }

	// closestHandle := gamemath.GetClosest(vector.NewWithValues([]float64{position.X, position.Y}), enemies)
	// if closestHandle == playerHandle {
	// 	return
	// }
	// closestEnemy := w.GetEntity(closestHandle)
	// launchProjectile(w, *position, *closestEnemy.GetPosition())
	// state.Set(component.AttackingState)
}

func launchProjectile(w world.World, from component.Position, to component.Position) {
	// 	e := gamemath.GetVector(from.X, from.Y)
	// 	m := gamemath.GetVector(to.X, to.Y)
	// 	_, entity := entity.Make(w, "entities/rocket.entity.json", from.X, from.Y)
	// 	direction := entity.GetDirection()
	// 	direction.Angle = gamemath.GetHeading(e, m)
}
