package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/filesystem"
	"cheezewiz/internal/input"
	"cheezewiz/internal/tag"
	"cheezewiz/internal/world"
	"cheezewiz/pkg/animation"
	"math/rand"
)

func Make(w world.World, path string, x float64, y float64) (world.EntityHandle, Entity) {
	var e EntityConfig

	e.Unmarshal(filesystem.Game.GetEntity(path))

	return w.Add(buildEntity(e, x, y))
}
func MakeWithTags(w world.World, path string, x float64, y float64, tags []tag.Tag) (world.EntityHandle, Entity) {
	var e EntityConfig

	e.Unmarshal(filesystem.Game.GetEntity(path))
	entity := buildEntity(e, x, y)
	for _, t := range tags {
		entity.AddTag(t)
	}

	return w.Add(entity)
}

func MakeRand(w world.World, path []string, x float64, y float64) (world.EntityHandle, Entity) {
	var e EntityConfig
	e.Unmarshal(filesystem.Game.GetEntity(path[rand.Intn(len(path))]))

	return w.Add(buildEntity(e, x, y))
}

func buildEntity(e EntityConfig, x float64, y float64) Entity {
	return entity{
		Position:    e.buildPosition(x, y),
		Health:      e.buildHealth(),
		Animation:   e.buildAnimations(),
		InputDevice: e.buildInputDevice(),
		State:       e.buildState(),
		RigidBody:   e.buildRigidBody(),
		TagSet:      e.buildTagSet(),
		Direction:   &component.Direction{},
	}
}

func lookupInputDevice(key string) input.PlayerInput {
	if key == "keyboard" {
		return input.Keyboard{}
	}
	return input.Keyboard{}
}

func (e EntityConfig) buildInputDevice() *component.InputDevice {
	if e.InputDevice == "" {
		return nil
	}
	return &component.InputDevice{
		Device: lookupInputDevice(e.InputDevice),
	}
}

func (e EntityConfig) buildHealth() *component.Health {
	return &e.Health
}
func (e EntityConfig) buildAnimations() *component.Animation {
	if e.Animations == nil {
		return nil
	}
	return &component.Animation{
		Animation: e.getAnimations(),
	}
}
func (e EntityConfig) getAnimations() map[component.StateType]*animation.Animation {
	anim := map[component.StateType]*animation.Animation{
		component.DebugState: animation.MakeDebugAnimation(),
	}

	for label, path := range e.Animations {
		anim[stringToState(label)] = animation.MakeAnimation(path, 32, 32, &filesystem.Game)
	}

	return anim
}

func stringToState(label string) component.StateType {
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

func (e EntityConfig) buildState() *component.State {
	s := &component.State{}
	available := map[component.StateType]component.StateType{}
	for key, value := range e.Animations {
		available[stringToState(key)] = stringToState(value)
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
	// ch := collision.HandlerLabel(e.RigidBody.CollisionHandlerLabel)
	// if ch == "" {
	// 	return &rb
	// }
	// rb.SetCollisionHandler(ch)
	return &rb
}

func tagLookup(t string) tag.Tag {
	switch t {
	case "player":
		return tag.Player
	case "actor":
		return tag.Actor
	case "collidable":
		return tag.Collidable
	case "enemy":
		return tag.Enemy
	case "projectile":
		return tag.Projectile
	case "animatable":
		return tag.Animatable
	case "viewport":
		return tag.ViewPort
	default:
		return tag.Nil
	}
}
func (e EntityConfig) buildTagSet() *component.TagSet {
	ts := component.NewTagSet()
	for _, t := range e.Tags {
		ts.Add(tagLookup(t))
	}
	return ts
}
