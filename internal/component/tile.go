package component

import "cheezewiz/pkg/ecs"

type Tile struct {
	// the actors that occupy this space
	Actors map[ecs.EntityHandle]ecs.Entity
}

func (t *Tile) Enter(handle ecs.EntityHandle, e ecs.Entity) {
	t.Actors[handle] = e
}

func (t *Tile) Exit(handle ecs.EntityHandle) {
	delete(t.Actors, handle)
}
