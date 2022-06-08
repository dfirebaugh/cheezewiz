package dentity

import (
	"cheezewiz/internal/collision"
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/internal/tag"
	"cheezewiz/pkg/animation"
	"math/rand"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type componentLabel string

// the elements in the components array should match these labels
const (
	PlayerTag     componentLabel = "playerTag"
	EnemyTag      componentLabel = "enemyTag"
	ProjectileTag componentLabel = "projectileTag"
	JellyBeanTag  componentLabel = "jellyBeanTag"
	XP            componentLabel = "xp"
	SpriteSheet   componentLabel = "spriteSheet"
	Position      componentLabel = "position"
	RigidBody     componentLabel = "rigidBody"
	Animations    componentLabel = "animations"
	Health        componentLabel = "health"
	Direction     componentLabel = "direction"
	InputDevice   componentLabel = "inputDevice"
	ActorState    componentLabel = "actorState"
)

var componentTable = map[componentLabel]*donburi.ComponentType{
	PlayerTag:     tag.Player,
	EnemyTag:      tag.Enemy,
	ProjectileTag: tag.Projectile,
	JellyBeanTag:  tag.JellyBean,
	XP:            component.XP,
	SpriteSheet:   component.SpriteSheet,
	Position:      component.Position,
	RigidBody:     component.RigidBody,
	Animations:    component.Animation,
	Health:        component.Health,
	Direction:     component.Direction,
	InputDevice:   component.InputDevice,
	ActorState:    component.ActorState,
}

// DynamicEntity is an entity that can be configured at runtime by parsing a json file
//  the structue of the json file will have to marshal out correctly
type DynamicEntity struct {
	PlayerTag     *donburi.ComponentType
	EnemyTag      *donburi.ComponentType
	ProjectileTag *donburi.ComponentType
	JellyBeanTag  *donburi.ComponentType
	config        EntityConfig
	XP            *component.XPData
	SpriteSheet   *component.SpriteSheetData
	Position      *component.PositionData
	RigidBody     *component.RigidBodyData
	Animation     *component.AnimationData
	ActorState    *component.ActorStateData
}

// MakeWithDirection entity must already have direction in it's config file
func MakeWithDirection(w donburi.World, path string, x float64, y float64, dir float64) *donburi.Entry {
	entry := MakeEntity(w, path, x, y)

	direction := component.GetDirection(entry)
	direction.Angle = dir

	return entry
}

func MakeRandEntity(w donburi.World, path []string, x float64, y float64) *donburi.Entry {
	var e EntityConfig
	var d DynamicEntity

	e.Unmarshal(pathToBytes(path[rand.Intn(len(path))]))
	d.config = e

	entry := w.Entry(w.Create(getComponents(e)...))

	err := d.initializeValues(entry, x, y)
	if err != nil {
		logrus.Error(err)
	}
	return entry
}

func MakeEntity(w donburi.World, path string, x float64, y float64) *donburi.Entry {
	var e EntityConfig
	var d DynamicEntity

	e.Unmarshal(pathToBytes(path))

	d.config = e

	entry := w.Entry(w.Create(getComponents(e)...))

	err := d.initializeValues(entry, x, y)
	if err != nil {
		logrus.Error(err)
	}
	return entry
}

// getComponents looks at the config and maps each componentlabel to a donburi.ComponentType
func getComponents(e EntityConfig) []*donburi.ComponentType {
	var ct []*donburi.ComponentType
	for _, key := range e.Components {
		if componentTable[key] == nil {
			logrus.Warn("component was not added: ", key, componentTable[key], componentTable[key] == nil)
			continue
		}
		ct = append(ct, componentTable[key])
	}
	return ct
}

func (d *DynamicEntity) initializeValues(entry *donburi.Entry, x, y float64) error {
	d.setPosition(entry, x, y)
	d.setXP(entry)
	d.setRigidBody(entry)
	d.setAnimations(entry)
	d.setActorState(entry)
	d.setInputDevice(entry)

	return nil
}

func (d *DynamicEntity) setXP(entry *donburi.Entry) {
	if !d.HasComponent(XP) {
		return
	}
	xp := component.GetXP(entry)
	xp.Value = d.config.XP
}
func (d *DynamicEntity) setPosition(entry *donburi.Entry, x, y float64) {
	if !d.HasComponent(Position) {
		return
	}
	position := component.GetPosition(entry)
	position.Set(x, y, d.config.Position.CX, d.config.Position.CY)
}
func (d *DynamicEntity) setRigidBody(entry *donburi.Entry) {
	if !d.HasComponent(RigidBody) {
		return
	}
	rb := component.GetRigidBody(entry)
	rb.SetBorder(d.config.RigidBody.R, d.config.RigidBody.B)
	ch := collision.HandlerLabel(d.config.RigidBody.CollisionHandlerLabel)
	rb.SetCollisionHandler(ch)
}

func (d *DynamicEntity) HasComponent(label componentLabel) bool {
	if len(d.config.Components) == 0 {
		return false
	}
	for _, value := range d.config.Components {
		if value != label {
			continue
		}

		return true
	}

	return false
}

func (d *DynamicEntity) setAnimations(entry *donburi.Entry) {
	if !d.HasComponent(Animations) {
		return
	}

	anim := component.GetAnimation(entry)
	anim.Animations = map[string]*animation.Animation{
		string(component.DebugState): animation.MakeDebugAnimation(),
	}

	for label, path := range d.config.Animations {
		anim.Animations[label] = animation.MakeAnimation(path, 32, 32)
	}
}
func (d *DynamicEntity) setActorState(entry *donburi.Entry) {
	if !d.HasComponent(ActorState) {
		return
	}
	s := component.GetActorState(entry)

	s.Available = map[string]string{}
	for label := range d.config.Animations {
		s.Available[label] = label
	}

	if d.config.ActorState == "" {
		s.Set(component.IdleState)
	}
	s.Set(component.ActorStateType(d.config.ActorState))
}
func (d *DynamicEntity) setInputDevice(entry *donburi.Entry) {
	if !d.HasComponent(InputDevice) {
		return
	}

	inputDevice := component.GetInputDevice(entry)

	if d.config.InputDevice == "keyboard" {
		inputDevice.Device = input.Keyboard{}
	}
}
