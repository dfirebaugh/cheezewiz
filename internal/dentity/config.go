package dentity

import (
	"cheezewiz/internal/component"
	"encoding/json"
)

type EntityConfig struct {
	Components  []componentLabel
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
