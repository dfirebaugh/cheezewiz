package scene

import (
	"cheezewiz/config/levels"
	"cheezewiz/internal/event"
	"encoding/json"
)

type worldDefinition struct {
	LevelMap  string             `json:"levelMap"`
	Countdown uint32             `json:"countdown"`
	Events    []event.SceneEvent `json:"events"`
	LevelName string             `json:"levelName"`
}

func (w *worldDefinition) Unmarshal(raw []byte) *worldDefinition {
	json.Unmarshal(raw, &w)

	return w
}

func LoadLevelOne() worldDefinition {
	w := worldDefinition{}
	w.Unmarshal(levels.Level1Raw)
	return w
}
func LoadStressTest() worldDefinition {
	w := worldDefinition{}
	w.Unmarshal(levels.StressTest)
	return w
}
