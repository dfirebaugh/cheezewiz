package entity

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"
)

func MakePlayer(w ecs.World) archetype.Player {
	a := archetype.Player{}
	a.ActorStateData = &component.ActorStateData{}
	a.AnimationData = &component.AnimationData{
		Animations: map[string]*animation.Animation{
			"debug": animation.MakeDebugAnimation(),
		},
	}
	a.InputDeviceData = &component.InputDeviceData{
		Device: input.Keyboard{},
	}
	a.PositionData = &component.PositionData{}

	w.Add(a)

	return a
}
