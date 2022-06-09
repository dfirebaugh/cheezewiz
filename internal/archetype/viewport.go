package archetype

import (
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"
)

type ViewPortTag ecs.Tag
type ViewPortArchetype struct {
	*component.PositionData
	ViewPortTag
}

func (v ViewPortArchetype) ViewPort() ViewPortTag {
	return v.ViewPortTag
}
func (v ViewPortArchetype) GetPosition() *component.PositionData {
	return v.PositionData
}
