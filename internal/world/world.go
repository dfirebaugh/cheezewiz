package world

import (
	"cheezewiz/internal/world/adapter"
	"sync"

	"github.com/google/uuid"
)

type World interface {
	Add(adapter.Entity) (EntityHandle, adapter.Entity)
	Remove(EntityHandle)
	// GetEntity - retrieve an entity by passing in it's handle
	GetEntity(EntityHandle) adapter.Entity
	// EachEntity - run a callback on each entity
	EachEntity(fn func(e EntityHandle))
	// FilterEntities - Filter Entities by a Predicate function
	FilterEntities(fn func(e EntityHandle) bool) []EntityHandle
	// First return the first entity handle that matches a condition
	First(fn func(handle EntityHandle) bool) EntityHandle
	// Len is a count of all entities in the world
	Len() int
	GetCacheIndex() int
}

type EntityHandle uuid.UUID

var nilEntityHandle = uuid.Nil

func (eh EntityHandle) IsNil() bool {
	return eh == EntityHandle(nilEntityHandle)
}

type world struct {
	entityMap map[EntityHandle]adapter.Entity
	mut       sync.RWMutex

	// the cacheIndex is incremented any time we add or remove entities
	cacheIndex int
}

func New() *world {
	return &world{
		entityMap:  map[EntityHandle]adapter.Entity{},
		cacheIndex: 0,
	}
}

func (w *world) Add(entity adapter.Entity) (EntityHandle, adapter.Entity) {
	w.mut.Lock()
	defer w.mut.Unlock()

	handle := EntityHandle(uuid.New())
	w.entityMap[handle] = entity
	w.cacheIndex++

	return handle, entity
}

func (w *world) Remove(handle EntityHandle) {
	w.mut.Lock()
	defer w.mut.Unlock()

	for h := range w.entityMap {
		if h == handle {
			delete(w.entityMap, handle)
		}
	}
	w.cacheIndex++
}

func (w *world) EachEntity(fn func(e EntityHandle)) {
	for h := range w.entityMap {
		fn(h)
	}
}

func (w *world) GetEntity(handle EntityHandle) adapter.Entity {
	return w.entityMap[handle]
}

func (w *world) FilterEntities(fn func(e EntityHandle) bool) []EntityHandle {
	handles := []EntityHandle{}
	w.EachEntity(func(handle EntityHandle) {
		if fn(handle) {
			handles = append(handles, handle)
		}
	})

	return handles
}
func (w *world) First(fn func(handle EntityHandle) bool) EntityHandle {
	h := EntityHandle(uuid.Nil)
	w.EachEntity(func(handle EntityHandle) {
		if fn(handle) {
			h = handle
		}
	})

	return h
}

func (w *world) Len() int {
	return len(w.entityMap)
}

func (w *world) GetCacheIndex() int {
	return w.cacheIndex
}
