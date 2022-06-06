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

func loadWorld(path string) worldDefinition {

	var l worldDefinition
	json.Unmarshal(levels.Level1Raw, &l)

	return l

}
