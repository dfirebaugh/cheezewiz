package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"
)

type ViewPort struct {
	*component.Position
	ViewPortTag ecs.Tag
}

func (v ViewPort) ViewPort() ecs.Tag {
	return v.ViewPortTag
}
func (v ViewPort) GetPosition() *component.Position {
	return v.Position
}
