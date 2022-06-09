package ecs

import (
	"errors"
	"sort"
)

type Entity interface{}
type Tag interface{}

type World struct {
	EntityMap map[int]Entity
}

func NewWorld() World {
	return World{
		EntityMap: map[int]Entity{},
	}
}

func NewTag() Tag {
	return struct{}{}
}

func (w *World) Add(entity Entity) (int, Entity) {
	uuid := len(w.EntityMap)
	w.EntityMap[uuid] = entity
	return uuid, entity
}

func (w World) GetEntities() map[int]Entity {
	return w.EntityMap
}

// getSortedEntityHandles returns the entities in order.
//   Otherwise, they would render in a seemingly random order.
func (w World) getSortedEntityHandles() []int {
	handles := make([]int, 0)
	for handle := range w.EntityMap {
		handles = append(handles, handle)
	}
	sort.Ints(handles)
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
