package adapter

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/ecs/cache"
	"cheezewiz/pkg/ecs"
)

type CustomECS struct {
	World ecs.World
}

func (cs CustomECS) GetWorld() ecs.World {
	return cs.World
}

func (cs CustomECS) Add(entity ecs.Entity) (ecs.EntityHandle, ecs.Entity) {
	return cs.GetWorld().Add(entity)
}

func (ce CustomECS) GetEnemies() ([]archetype.Enemy, bool) {
	return cache.GetEnemies(ce.World)
}
func (ce CustomECS) GetAnimatables() ([]archetype.Animatable, bool) {
	return cache.GetAnimatables(ce.World)
}
func (ce CustomECS) GetCollidables() ([]archetype.Collidable, bool) {
	return cache.GetCollidables(ce.World)
}
func (ce CustomECS) GetProjectiles() ([]archetype.Projectile, bool) {
	return cache.GetProjectiles(ce.World)
}
func (ce CustomECS) GetPlayers() ([]archetype.Player, bool) {
	return cache.GetPlayers(ce.World)
}
func (ce CustomECS) GetPlayerHandles() ([]ecs.EntityHandle, bool) {
	return cache.GetPlayerHandles(ce.World)
}
func (ce CustomECS) GetEntity(handle ecs.EntityHandle) ecs.Entity {
	return ce.World.GetEntity(handle)
}
func (ce CustomECS) FirstViewPort() (archetype.ViewPort, error) {
	return ecs.FirstEntity[archetype.ViewPort](ce.World)
}
func (ce CustomECS) FirstPlayer() (archetype.Player, error) {
	return ecs.FirstEntity[archetype.Player](ce.World)
}
func (ce CustomECS) Count() int {
	return ce.World.Count()
}
func (ce CustomECS) Remove(e ecs.Entity) {
	ce.World.Remove(e)
}
