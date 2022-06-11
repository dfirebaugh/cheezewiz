package adapter

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/pkg/ecs"
)

type Adapter interface {
	GetWorld() ecs.World
	Add(entity ecs.Entity) (ecs.EntityHandle, ecs.Entity)
	GetEntity(handle ecs.EntityHandle) ecs.Entity
	GetEnemies() ([]archetype.Enemy, bool)
	GetProjectiles() ([]archetype.Projectile, bool)
	GetPlayers() ([]archetype.Player, bool)
	GetPlayerHandles() ([]ecs.EntityHandle, bool)
	// GetActors() ([]archetype.Collidable, bool)
	GetCollidables() ([]archetype.Collidable, bool)
	GetAnimatables() ([]archetype.Animatable, bool)
	FirstViewPort() (archetype.ViewPort, error)
	FirstPlayer() (archetype.Player, error)
	Remove(ecs.Entity)
	Count() int
}

func Adapt(world interface{}) Adapter {
	return CustomECS{
		World: world.(ecs.World),
	}
}
