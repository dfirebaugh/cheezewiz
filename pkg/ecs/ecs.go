package ecs

import (
	"errors"
	"sort"

	"github.com/google/uuid"
)

type Entity interface{}
type Tag interface{}

type World struct {
	EntityMap map[uuid.UUID]Entity
}

func NewWorld() World {
	return World{
		EntityMap: map[uuid.UUID]Entity{},
	}
}

func NewTag() Tag {
	return struct{}{}
}

func (w *World) Add(entity Entity) (uuid.UUID, Entity) {
	uuid := uuid.New()
	w.EntityMap[uuid] = entity
	return uuid, entity
}

func (w World) GetEntities() map[uuid.UUID]Entity {
	return w.EntityMap
}

// getSortedEntityHandles returns the entities in order.
//   Otherwise, they would render in a seemingly random order.
func (w World) getSortedEntityHandles() []uuid.UUID {
	handles := make([]uuid.UUID, 0)
	for handle := range w.EntityMap {
		handles = append(handles, handle)
	}

	sort.SliceStable(handles, func(i, j int) bool {
		return handles[i].ID() < handles[j].ID()
	})

	return handles
}

func FilterBy[T any](w World) []T {
	var result []T
	for _, h := range w.getSortedEntityHandles() {
		if e, ok := w.EntityMap[h].(T); ok {
			result = append(result, e)
		}
	}

	return result
}

// returns an entity that matches the filter or nil
func FirstEntity[T any](w World) (T, error) {
	filtered := FilterBy[T](w)
	l := len(filtered)
	var first T
	if l == 0 {
		return first, errors.New("unable to find an entity of that type")
	}
	return filtered[0], nil
}

func IsType[T any](entity Entity) bool {
	_, ok := entity.(T)
	return ok
}
