package entity

import (
	"cheezewiz/internal/component"
	"encoding/json"
)

// DynamicEntity is an entity that can be configured at runtime by parsing a json file
//  the structue of the json file will have to marshal out correctly
type DynamicEntity struct {
	config      EntityConfig
	XP          *component.XP
	SpriteSheet *component.SpriteSheet
	Position    *component.Position
	RigidBody   *component.RigidBody
	Animation   *component.Animation
	State       *component.State
}

type EntityConfig struct {
	Archetype   string              `json:"archetype"`
	XP          float64             `json:"xp"`
	SpriteSheet string              `json:"spriteSheet"`
	Position    component.Position  `json:"position"`
	RigidBody   component.RigidBody `json:"rigidBody"`
	Direction   component.Direction `json:"direction"`
	Health      component.Health    `json:"health"`
	Animations  map[string]string   `json:"animations"`
	State       string              `json:"state"`
	InputDevice string              `json:"inputDevice"`
}

func (e *EntityConfig) Unmarshal(bytes []byte) {
	json.Unmarshal(bytes, e)
}
