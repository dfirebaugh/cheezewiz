package filter

import (
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query/predicate"
)

func filter(handles []world.EntityHandle, p predicate.EntityPredicate) []world.EntityHandle {
	result := []world.EntityHandle{}
	for _, handle := range handles {
		if !p(handle) {
			continue
		}

		result = append(result, handle)
	}

	return result
}

func GetPlayers(w world.World, h []world.EntityHandle) []world.EntityHandle {
	return filter(h, predicate.IsPlayer(w))
}
func GetProjectiles(w world.World, h []world.EntityHandle) []world.EntityHandle {
	return filter(h, predicate.IsProjectile(w))
}
func GetAnimatables(w world.World, h []world.EntityHandle) []world.EntityHandle {
	return filter(h, predicate.IsAnimatable(w))
}
func GetEnemies(w world.World, h []world.EntityHandle) []world.EntityHandle {
	return filter(h, predicate.IsEnemy(w))
}
