package ecs

import (
	"errors"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type Entity interface{}
type Tag interface{}

type World interface {
	Add(entity Entity) (uuid.UUID, Entity)
	GetEntities() map[uuid.UUID]Entity
	GetSortedEntityHandles() []uuid.UUID
	GetEntity(id uuid.UUID) Entity
}

type world struct {
	EntityMap map[uuid.UUID]Entity
	mut       sync.RWMutex
}

func NewWorld() *world {
	return &world{
		EntityMap: map[uuid.UUID]Entity{},
		mut:       sync.RWMutex{},
	}
}

func NewTag() Tag {
	return struct{}{}
}

func (w *world) Add(entity Entity) (uuid.UUID, Entity) {
	w.mut.Lock()
	defer w.mut.Unlock()

	uuid := uuid.New()
	w.EntityMap[uuid] = entity
	return uuid, entity
}

func (w *world) GetEntities() map[uuid.UUID]Entity {
	return w.EntityMap
}

func (w *world) GetEntity(id uuid.UUID) Entity {
	w.mut.Lock()
	defer w.mut.Unlock()
	return w.EntityMap[id]
}

// getSortedEntityHandles returns the entities in order.
//   Otherwise, they would render in a seemingly random order.
func (w *world) GetSortedEntityHandles() []uuid.UUID {
	w.mut.Lock()
	defer w.mut.Unlock()

	handles := make([]uuid.UUID, 0)
	for handle := range w.EntityMap {
		handles = append(handles, handle)
	}

	sort.SliceStable(handles, func(i, j int) bool {
		return handles[i].ID() < handles[j].ID()
	})

	return handles
}

// FilterMapBy is unordered, but allows you to get the
//  handle and entity
func FilterMapBy[T any](w World) map[uuid.UUID]T {
	result := map[uuid.UUID]T{}
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
