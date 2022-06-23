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

var CheeseMissile = func() func() {
	return func() {
		world.Instance.EachEntity(func(handle world.EntityHandle) {
			if !world.Instance.GetEntity(handle).HasTag(tag.Player) {
				return
			}
			player := world.Instance.GetEntity(handle)

			findHeading(player, handle)
		})
	}
}

func findHeading(player entity.Entity, playerHandle world.EntityHandle) {
	// position := player.GetPosition()
	// state := player.GetState()
	// if state.GetCurrent() == component.DeathState {
	// 	return
	// }

	// enemies := map[world.EntityHandle]vector.Vector{}

	// world.Instance.EachEntity(func(handle world.EntityHandle) {
	// 	if handle == playerHandle {
	// 		return
	// 	}
	// 	p := world.Instance.GetEntity(handle).GetPosition()
	// 	enemies[handle] = vector.NewWithValues([]float64{p.X, p.Y})
	// })

	// if len(enemies) == 0 {
	// 	return
	// }

	// closestHandle := gamemath.GetClosest(vector.NewWithValues([]float64{position.X, position.Y}), enemies)
	// if closestHandle == playerHandle {
	// 	return
	// }
	// closestEnemy := world.Instance.GetEntity(closestHandle)
	// launchProjectile(w, *position, *closestEnemy.GetPosition())
	// state.Set(component.AttackingState)
}

func launchProjectile(from component.Position, to component.Position) {
	// 	e := gamemath.GetVector(from.X, from.Y)
	// 	m := gamemath.GetVector(to.X, to.Y)
	// 	_, entity := entity.Make(world.Instance, "entities/rocket.entity.json", from.X, from.Y)
	// 	direction := entity.GetDirection()
	// 	direction.Angle = gamemath.GetHeading(e, m)
}
