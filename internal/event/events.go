package event

import (
	"github.com/yohamta/donburi"
)

type Job struct {
	json_name string
	callback  func(w donburi.World, args []string)
	priority  int
}

func spawnWave(w donburi.World, args []string) {

}

type JobName string

const (
	SpawnWave JobName = "SpawnWave"
	SpawnBoss JobName = "SpawnBoss"
	HurryUp   JobName = "HurryUp"
)

var JobTypes = make(map[JobName]Job, 1)

type SceneEvent struct {
	Time      uint32   `json:"time"`
	EventName string   `json:"event"`
	EventArgs []string `json:"args"`
}

func init() {
	JobTypes[SpawnWave] = Job{
		json_name: "SpawnWave",
		callback:  spawnWave,
		priority:  0,
	}
}
