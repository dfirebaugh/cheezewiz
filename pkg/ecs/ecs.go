package ecs

import (
	"errors"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type EntityHandle uuid.UUID
type Entity interface{}
type Tag interface{}

var NilEntityHandle = EntityHandle(uuid.Nil)

type World interface {
	Add(entity Entity) (EntityHandle, Entity)
	GetEntities() map[EntityHandle]Entity
	GetSortedEntityHandles() []EntityHandle
	GetEntity(id EntityHandle) Entity
}

type world struct {
	EntityMap map[EntityHandle]Entity
	mut       sync.RWMutex
}

func NewWorld() *world {
	return &world{
		EntityMap: map[EntityHandle]Entity{},
		mut:       sync.RWMutex{},
	}
}

func NewTag() Tag {
	return struct{}{}
}

func NewEntityHandle() EntityHandle {
	return EntityHandle(uuid.New())
}

func (w *world) Add(entity Entity) (EntityHandle, Entity) {
	w.mut.Lock()
	defer w.mut.Unlock()

	id := NewEntityHandle()
	w.EntityMap[id] = entity
	return id, entity
}

func (w *world) GetEntities() map[EntityHandle]Entity {
	return w.EntityMap
}

func (w *world) GetEntity(id EntityHandle) Entity {
	w.mut.Lock()
	defer w.mut.Unlock()
	return w.EntityMap[id]
}

// getSortedEntityHandles returns the entities in order.
//   Otherwise, they would render in a seemingly random order.
func (w *world) GetSortedEntityHandles() []EntityHandle {
	w.mut.Lock()
	defer w.mut.Unlock()

	handles := make([]EntityHandle, 0)
	for handle := range w.EntityMap {
		handles = append(handles, handle)
	}

	sort.SliceStable(handles, func(i, j int) bool {
		return uuid.UUID(handles[i]).ID() < uuid.UUID(handles[j]).ID()
	})

	return handles
}

// FilterMapBy is unordered, but allows you to get the
//  handle and entity
func FilterMapBy[T any](w World) map[EntityHandle]T {
	result := map[EntityHandle]T{}
	for handle, e := range w.GetEntities() {
		if e, ok := e.(T); ok {
			result[handle] = e
		}
	}

	return result
}

func FilterBy[T any](w World) []T {
	var result []T
	for _, h := range w.GetSortedEntityHandles() {
		if e, ok := w.GetEntity(h).(T); ok {
			result = append(result, e)
		}
	}

	return result
}

// returns an entity that matches the filter or nil
func FirstEntity[T any](w World) (T, error) {
	var first T
	filtered := FilterBy[T](w)
	l := len(filtered)
	if l == 0 {
		return first, errors.New("unable to find an entity of that type")
	}
	return filtered[0], nil
}

func IsType[T any](entity Entity) bool {
	_, ok := entity.(T)
	return ok
}
