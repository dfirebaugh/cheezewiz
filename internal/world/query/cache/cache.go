package cache

import (
	"cheezewiz/internal/world"

	cache "github.com/Code-Hex/go-generics-cache"
)

var entityCache = cache.New[uint, []world.EntityHandle]()

const cacheKey = 1

var prevCacheIndex int

func Get(w world.World) ([]world.EntityHandle, bool) {
	refreshEntities(w)

	return entityCache.Get(cacheKey)
}

func refreshEntities(w world.World) {
	if !shouldRefresh(w) {
		return
	}
	entityCache.Set(cacheKey, getLatest(w))
}

func shouldRefresh(w world.World) bool {
	latest := w.GetCacheIndex()
	defer updateCacheIndex(latest)

	return prevCacheIndex != latest
}

func updateCacheIndex(new int) {
	prevCacheIndex = new
}

func getLatest(w world.World) []world.EntityHandle {
	e := []world.EntityHandle{}
	w.EachEntity(func(h world.EntityHandle) {
		e = append(e, h)
	})

	return e
}
