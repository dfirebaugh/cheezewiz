package ecs

type Entity interface{}
type FilterID int

type filterFn func(entity Entity) bool

type World struct {
	EntityMap map[int]Entity
	// might not need the slice of handlers
	Entities []int

	filterCache map[FilterID]filterFn
}

func NewWorld() World {
	return World{
		EntityMap:   map[int]Entity{},
		filterCache: map[FilterID]filterFn{},
	}
}

func (w *World) Add(entity Entity) (int, Entity) {
	uuid := len(w.EntityMap)
	w.Entities = append(w.Entities, uuid)
	w.EntityMap[uuid] = entity
	return uuid, entity
}

func (w World) GetEntities() map[int]Entity {
	return w.EntityMap
}

func (w *World) MakeFilter(filter filterFn) FilterID {
	var fID FilterID = FilterID(len(w.filterCache))
	w.filterCache[fID] = filter
	return fID
}

func (w World) FilterBy(fID FilterID) []Entity {
	var result []Entity
	for _, ent := range w.EntityMap {
		if w.filterCache[fID](ent) {
			result = append(result, ent)
		}
	}

	return result
}

// returns an entity that matches the filter or nil
func (w World) FirstEntity(fID FilterID) Entity {
	for _, e := range w.FilterBy(fID) {
		return e
	}
	return nil
}
