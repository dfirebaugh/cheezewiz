package query

import (
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query/cache"
)

type filter func(w world.World, h []world.EntityHandle) []world.EntityHandle
type entityiterator func(handle world.EntityHandle)

// Each, filter out entities in the world based on a defined predicate and iterate
//   through them to preform some action
func Each(w world.World, f filter, iter entityiterator) {
	handles, _ := cache.Get(w)

	for _, handle := range f(w, handles) {
		iter(handle)
	}
}
