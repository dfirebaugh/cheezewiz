package entity

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/internal/filesystem"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"
	"math/rand"
)

func MakeEntity(w ecs.World, path string, x float64, y float64) (int, interface{}) {
	var e EntityConfig

	e.Unmarshal(filesystem.GetEntity(path))

	return w.Add(buildArchetype(e))
}
func MakeRandEntity(w ecs.World, path []string, x float64, y float64) (int, interface{}) {
	var e EntityConfig

	e.Unmarshal(filesystem.GetEntity(path[rand.Intn(len(path))]))

	return w.Add(buildArchetype(e))
}

func buildArchetype(e EntityConfig) interface{} {
	// check entities componentlabels and build archetype based on that?
	switch e.Archetype {
	case "player":
		return buildPlayer(e)
	case "actor":
		return buildActor(e)
	case "enemy":
		return buildActor(e)
	}

	return nil
}

func buildPlayer(e EntityConfig) *archetype.Player {
	p := archetype.Player{
		PositionData: &e.Position,
		HealthAspect: &e.Health,
		AnimationData: &component.AnimationData{
			Animations: e.getAnimations(),
		},
		InputDeviceData: &component.InputDeviceData{
			Device: lookupInputDevice(e.InputDevice),
		},
		ActorStateData: e.getState(),
	}
	return &p
}
func buildActor(e EntityConfig) *archetype.Actor {
	p := archetype.Actor{
		PositionData: &e.Position,
		HealthAspect: &e.Health,
		AnimationData: &component.AnimationData{
			Animations: e.getAnimations(),
		},
		ActorStateData: e.getState(),
	}
	return &p
}

func lookupInputDevice(key string) input.PlayerInput {
	if key == "keyboard" {
		return input.Keyboard{}
	}
	return input.Keyboard{}
}

func (e EntityConfig) getAnimations() map[string]*animation.Animation {
	anim := map[string]*animation.Animation{
		string(component.DebugState): animation.MakeDebugAnimation(),
	}

	for label, path := range e.Animations {
		anim[label] = animation.MakeAnimation(path, 32, 32)
	}

	return anim
}

func (e EntityConfig) getState() *component.ActorStateData {
	s := &component.ActorStateData{}
	s.SetAvailable(e.Animations)
	s.Set(component.ActorStateType(e.ActorState))
	return s
}

func (e EntityConfig) HasComponent(label componentLabel) bool {
	if len(e.Components) == 0 {
		return false
	}
	for _, value := range e.Components {
		if value != label {
			continue
		}

		return true
	}

	return false
}
