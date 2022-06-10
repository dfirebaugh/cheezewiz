package cache

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/pkg/ecs"

	cache "github.com/Code-Hex/go-generics-cache"
)

type cacheKey uint

const (
	animatableKey cacheKey = iota
	collidableKey
	playerKey
	enemyKey
)

var player = cache.New[cacheKey, []archetype.Player]()
var enemy = cache.New[cacheKey, []archetype.Enemy]()
var animatable = cache.New[cacheKey, []archetype.Animatable]()
var collidable = cache.New[cacheKey, []archetype.Collidable]()

var prevEntityCount int

func GetEnemies(w ecs.World) ([]archetype.Enemy, bool) {
	refreshEnemies(w)
	prevEntityCount = w.Count()
	return enemy.Get(enemyKey)
}
func GetAnimatables(w ecs.World) ([]archetype.Animatable, bool) {
	refreshAnimatables(w)
	prevEntityCount = w.Count()
	return animatable.Get(animatableKey)
}
func GetCollidables(w ecs.World) ([]archetype.Collidable, bool) {
	refreshCollidables(w)
	prevEntityCount = w.Count()
	return collidable.Get(collidableKey)
}
func GetPlayers(w ecs.World) ([]archetype.Player, bool) {
	refreshPlayers(w)
	prevEntityCount = w.Count()
	return player.Get(playerKey)
}

func RefreshCache(w ecs.World) {
	refreshEnemies(w)
	refreshPlayers(w)
	refreshAnimatables(w)
	prevEntityCount = w.Count()
}

func refreshEnemies(w ecs.World) {
	if !enemyShouldRefresh(w) {
		return
	}
	enemy.Set(enemyKey, ecs.FilterBy[archetype.Enemy](w))
}
func refreshPlayers(w ecs.World) {
	if !playerShouldRefresh(w) {
		return
	}
	player.Set(playerKey, ecs.FilterBy[archetype.Player](w))
}
func refreshAnimatables(w ecs.World) {
	if !animatableShouldRefresh(w) {
		return
	}
	println("update animatables")
	animatable.Set(animatableKey, ecs.FilterBySorted[archetype.Animatable](w))
}
func refreshCollidables(w ecs.World) {
	if !collidablesShouldRefresh(w) {
		return
	}
	collidable.Set(collidableKey, ecs.FilterBy[archetype.Collidable](w))
}

func collidablesShouldRefresh(w ecs.World) bool {
	if len(collidable.Keys()) == 0 {
		return true
	}
	return prevEntityCount == w.Count()
}

func enemyShouldRefresh(w ecs.World) bool {
	if len(enemy.Keys()) == 0 {
		return true
	}
	return prevEntityCount == w.Count()
}
func animatableShouldRefresh(w ecs.World) bool {
	if len(animatable.Keys()) == 0 {
		return true
	}
	return prevEntityCount == w.Count()
}
func playerShouldRefresh(w ecs.World) bool {
	return true
}
