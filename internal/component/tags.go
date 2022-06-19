package component

import "cheezewiz/internal/tag"

type TagSet struct {
	Tags map[tag.Tag]struct{}
}

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
)

var exists = struct{}{}

func NewTagSet() *TagSet {
	return &TagSet{
		Tags: make(map[tag.Tag]struct{}),
	}
}
func (ts *TagSet) Add(t tag.Tag) {
	ts.Tags[t] = exists
}

func (ts *TagSet) Contains(t tag.Tag) bool {
	_, ok := ts.Tags[t]
	return ok
}
