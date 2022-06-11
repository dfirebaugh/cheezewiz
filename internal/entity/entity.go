package entity

import (
	"cheezewiz/internal/archetype"
	"cheezewiz/internal/collision"
	"cheezewiz/internal/component"
	"cheezewiz/internal/ecs/adapter"
	"cheezewiz/internal/filesystem"
	"cheezewiz/internal/input"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/ecs"
	"math/rand"

	"github.com/sirupsen/logrus"
)

type Directionable interface {
	GetDirection() *component.Direction
}

func MakeWithDirection(w adapter.Adapter, path string, x float64, y float64, dir float64) (ecs.EntityHandle, ecs.Entity) {
	handle, entity := MakeEntity(w, path, x, y)

	e, ok := entity.(Directionable)
	if !ok {
		logrus.Error("not able to get a direction for entity")
	}
	direction := e.GetDirection()
	if direction == nil {
		logrus.Error("not a valid direction")
		return ecs.NilEntityHandle, nil
	}

	direction.Angle = dir
	return handle, e
}

func MakeEntity(w adapter.Adapter, path string, x float64, y float64) (ecs.EntityHandle, ecs.Entity) {
	var e EntityConfig

	e.Unmarshal(filesystem.Game.GetEntity(path))

	return w.Add(buildEntity(e, x, y))
}
func MakeRandEntity(w adapter.Adapter, path []string, x float64, y float64) (ecs.EntityHandle, ecs.Entity) {
	var e EntityConfig

	e.Unmarshal(filesystem.Game.GetEntity(path[rand.Intn(len(path))]))

	return w.Add(buildEntity(e, x, y))
}

func buildEntity(e EntityConfig, x float64, y float64) (ecs.Entity, []ecs.ArchetypeTag) {
	// check entities componentlabels and build archetype based on that?
	switch e.Archetype {
	case "player":
		return buildPlayer(e, x, y)
	case "actor":
		return buildActor(e, x, y)
	case "enemy":
		return buildEnemy(e, x, y)
	case "projectile":
		return buildProjectile(e, x, y)
	default:
		logrus.Errorf("archetype is not defined: %s", e.Archetype)
	}

	return nil, nil
}

func buildPlayer(e EntityConfig, x float64, y float64) (*Player, []ecs.ArchetypeTag) {
	archetypes := []ecs.ArchetypeTag{ecs.ArchetypeTag(archetype.PlayerTag), ecs.ArchetypeTag(archetype.AnimatableTag)}
	p := Player{
		Position: e.buildPosition(x, y),
		Health:   &e.Health,
		Animation: &component.Animation{
			Animation: e.getAnimations(),
		},
		InputDevice: &component.InputDevice{
			Device: lookupInputDevice(e.InputDevice),
		},
		State:     e.getState(),
		RigidBody: e.buildRigidBody(),
	}
	return &p, archetypes
}
func buildProjectile(e EntityConfig, x float64, y float64) (*Projectile, []ecs.ArchetypeTag) {
	archetypes := []ecs.ArchetypeTag{ecs.ArchetypeTag(archetype.ProjectileTag), ecs.ArchetypeTag(archetype.AnimatableTag)}

	p := &Projectile{
		Position: e.buildPosition(x, y),
		Animation: &component.Animation{
			Animation: e.getAnimations(),
		},
		State:     e.getState(),
		RigidBody: e.buildRigidBody(),
		Direction: &component.Direction{},
	}
	return p, archetypes
}

func buildEnemy(e EntityConfig, x float64, y float64) (*Enemy, []ecs.ArchetypeTag) {
	archetypes := []ecs.ArchetypeTag{ecs.ArchetypeTag(archetype.EnemyTag), ecs.ArchetypeTag(archetype.AnimatableTag)}

	p := Enemy{
		Position: e.buildPosition(x, y),
		Health:   &e.Health,
		Animation: &component.Animation{
			Animation: e.getAnimations(),
		},
		State:     e.getState(),
		RigidBody: e.buildRigidBody(),
	}
	return &p, archetypes
}

func buildActor(e EntityConfig, x float64, y float64) (*Actor, []ecs.ArchetypeTag) {
	archetypes := []ecs.ArchetypeTag{ecs.ArchetypeTag(archetype.ActorTag), ecs.ArchetypeTag(archetype.AnimatableTag)}
	a := &Actor{
		Position: e.buildPosition(x, y),
		Health:   &e.Health,
		Animation: &component.Animation{
			Animation: e.getAnimations(),
		},
		State:     e.getState(),
		RigidBody: e.buildRigidBody(),
	}
	return a, archetypes
}

func lookupInputDevice(key string) input.PlayerInput {
	if key == "keyboard" {
		return input.Keyboard{}
	}
	return input.Keyboard{}
}

func (e EntityConfig) getAnimations() map[component.StateType]*animation.Animation {
	anim := map[component.StateType]*animation.Animation{
		component.DebugState: animation.MakeDebugAnimation(),
	}

	for label, path := range e.Animations {
		anim[e.stringToState(label)] = animation.MakeAnimation(path, 32, 32, &filesystem.Game)
	}

	return anim
}

func (e EntityConfig) stringToState(label string) component.StateType {
	switch label {
	case "debug":
		return component.DebugState
	case "idle":
		return component.IdleState
	case "walk":
		return component.WalkingState
	case "attack":
		return component.AttackingState
	case "hurt":
		return component.HurtState
	case "death":
		return component.DeathState
	default:
		return component.DebugState
	}
}

func (e EntityConfig) getState() *component.State {
	s := &component.State{}
	available := map[component.StateType]component.StateType{}
	for key, value := range e.Animations {
		available[e.stringToState(key)] = e.stringToState(value)
	}
	s.SetAvailable(available)
	s.Set(component.StateType(e.State))
	return s
}

func (e EntityConfig) buildPosition(x float64, y float64) *component.Position {
	position := &component.Position{
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
func (e EntityConfig) buildRigidBody() *component.RigidBody {
	rb := component.RigidBody{}
	rb.SetBorder(e.RigidBody.R, e.RigidBody.B)
	ch := collision.HandlerLabel(e.RigidBody.CollisionHandlerLabel)
	if ch == "" {
		return &rb
	}
	rb.SetCollisionHandler(ch)
	return &rb
}
