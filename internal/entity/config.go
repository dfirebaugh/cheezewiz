package entity

import (
	"cheezewiz/internal/component"
	"encoding/json"
)

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
	Tags        []string            `json:"tags"`
}

func (e *EntityConfig) Unmarshal(bytes []byte) {
	json.Unmarshal(bytes, e)
}
