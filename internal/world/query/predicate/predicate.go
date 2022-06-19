package predicate

import (
	"cheezewiz/internal/tag"
	"cheezewiz/internal/world"
)

type EntityPredicate func(handle world.EntityHandle) bool

func HasTag(w world.World, t tag.Tag) EntityPredicate {
	return func(handle world.EntityHandle) bool {
		return w.GetEntity(handle).HasTag(t)
	}
}
func IsPlayer(w world.World) EntityPredicate {
	return HasTag(w, tag.Player)
}
func IsProjectile(w world.World) EntityPredicate {
	return HasTag(w, tag.Projectile)
}
func IsAnimatable(w world.World) EntityPredicate {
	return HasTag(w, tag.Animatable)
}
func IsEnemy(w world.World) EntityPredicate {
	return HasTag(w, tag.Enemy)
}
