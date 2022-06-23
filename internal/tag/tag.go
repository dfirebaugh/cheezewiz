package tag

type Tag int

const (
	Nil Tag = iota
	Player
	Actor
	Collidable
	Enemy
	Projectile
	Animatable
	ViewPort
	Bound
)
