package system

import (
	"cheezewiz/examples/choppa/internal/component"
	"cheezewiz/examples/choppa/internal/entity"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type LifeSpan struct {
	projectileQuery *query.Query
}

func NewLifeSpan() *LifeSpan {
	return &LifeSpan{
		projectileQuery: query.NewQuery(filter.Contains(entity.ProjectileTag)),
	}
}

func (l LifeSpan) Update(w donburi.World) {
	// currently disabled because of an index out of range error
	// so, projectiles will not get destroyed until this is fixed
	disabled := true
	if disabled {
		return
	}
	l.projectileLife(w)
}

func (l LifeSpan) projectileLife(w donburi.World) {
	l.projectileQuery.EachEntity(w, func(entry *donburi.Entry) {
		tick := component.GetTick(entry)
		tick.Value++

		if tick.Value == tick.EOL {
			w.Remove(entry.Entity())
		}
	})
}
