package scene

import (
	"cheezewiz/internal/event"
	"cheezewiz/level"
	"encoding/json"
)

type worldDefinition struct {
	LevelMap  string             `json:"levelMap"`
	Countdown uint32             `json:"countdown"`
	Events    []event.SceneEvent `json:"events"`
	LevelName string             `json:"levelName"`
}

func loadWorld(path string) worldDefinition {

	var l worldDefinition
	json.Unmarshal(level.Level1Raw, &l)

	return l

}
