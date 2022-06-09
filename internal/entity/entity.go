package entity

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/component"
	"cheezewiz/internal/filesystem"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"
	"math/rand"

	"github.com/sirupsen/logrus"
)

func MakeEntity(w ecs.World, path string, x float64, y float64) (int, ecs.Entity) {
	var e EntityConfig

	e.Unmarshal(filesystem.GetEntity(path))

	return w.Add(buildArchetype(e, x, y))
}
func MakeRandEntity(w ecs.World, path []string, x float64, y float64) (int, ecs.Entity) {
	var e EntityConfig

	e.Unmarshal(filesystem.GetEntity(path[rand.Intn(len(path))]))

	return w.Add(buildArchetype(e, x, y))
}

func buildArchetype(e EntityConfig, x float64, y float64) ecs.Entity {
	// check entities componentlabels and build archetype based on that?
	switch e.Archetype {
	case "player":
		return buildPlayer(e, x, y)
	case "actor":
		return buildActor(e, x, y)
	case "enemy":
		return buildEnemy(e, x, y)
	default:
		logrus.Errorf("archetype is not defined: %s", e.Archetype)
	}

	return nil
}

func buildPlayer(e EntityConfig, x float64, y float64) *archetype.PlayerArchetype {
	p := archetype.PlayerArchetype{
		PositionData: e.buildPosition(x, y),
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
func buildActor(e EntityConfig, x float64, y float64) *archetype.ActorArchetype {
	p := archetype.ActorArchetype{
		PositionData: e.buildPosition(x, y),
		HealthAspect: &e.Health,
		AnimationData: &component.AnimationData{
			Animations: e.getAnimations(),
		},
		ActorStateData: e.getState(),
	}
	return &p
}
func buildEnemy(e EntityConfig, x float64, y float64) *archetype.EnemyArchetype {
	p := archetype.EnemyArchetype{
		PositionData: e.buildPosition(x, y),
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

func (e EntityConfig) buildPosition(x float64, y float64) *component.PositionData {
	position := &component.PositionData{
		X:  x,
		Y:  y,
		CX: e.Position.CX,
		CY: e.Position.CY,
	}
	if e.Position.X > 0 {
		position.X = e.Position.X
	}

	if e.Position.Y > 0 {
		position.Y = e.Position.Y
	}
	return position
}
