package scene

import (
	"cheezewiz/internal/event"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type worldDefinition struct {
	LevelMap  string             `json:"levelMap"`
	Countdown uint32             `json:"countdown"`
	Events    []event.SceneEvent `json:"events"`
	LevelName string             `json:"levelName"`
}

func loadWorld(path string) worldDefinition {

	var level worldDefinition
	var rawData, err = os.ReadFile(path)

	if err != nil {
		logrus.Fatal(err)
	}

	json.Unmarshal(rawData, &level)
	fmt.Printf("%#v", level)

	return level

}
