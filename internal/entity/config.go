package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/tag"
	"encoding/json"

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
	// PlayerTag     *donburi.ComponentType
	// EnemyTag      *donburi.ComponentType
	// ProjectileTag *donburi.ComponentType
	// JellyBeanTag  *donburi.ComponentType
	config      EntityConfig
	XP          *component.XPData
	SpriteSheet *component.SpriteSheetData
	Position    *component.PositionData
	RigidBody   *component.RigidBodyData
	Animation   *component.AnimationData
	ActorState  *component.ActorStateData
}

type EntityConfig struct {
	Components  []componentLabel
	Archetype   string                  `json:"archetype"`
	XP          float64                 `json:"xp"`
	SpriteSheet string                  `json:"spriteSheet"`
	Position    component.PositionData  `json:"position"`
	RigidBody   component.RigidBodyData `json:"rigidBody"`
	Direction   component.DirectionData `json:"direction"`
	Health      component.HealthAspect  `json:"health"`
	Animations  map[string]string       `json:"animations"`
	ActorState  string                  `json:"actorState"`
	InputDevice string                  `json:"inputDevice"`
}

func (e *EntityConfig) Unmarshal(bytes []byte) {
	json.Unmarshal(bytes, e)
	e.Components = getComponentLabels(bytes)
}

func getComponentLabels(rawJSON []byte) []componentLabel {
	var componentLabels []componentLabel
	var keys map[string]interface{}
	json.Unmarshal(rawJSON, &keys)

	for key := range keys {
		// take the keys of the json object and convert them to component labels
		label := componentLabel(key)
		componentLabels = append(componentLabels, label)
	}
	return componentLabels
}
