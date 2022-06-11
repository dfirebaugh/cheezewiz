package event

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/ecs/adapter"
	"cheezewiz/internal/entity"
	"cheezewiz/pkg/ecs"
	"math"
	"strconv"

	"github.com/atedja/go-vector"
	"github.com/sirupsen/logrus"
)

type Player interface {
	GetPosition() *component.Position
	GetPlayerTag() ecs.Tag
}

type Enemy interface {
	GetPosition() *component.Position
	GetHealth() *component.Health
	GetEnemyTag() ecs.Tag
}

type Job struct {
	json_name string
	Callback  func(w adapter.Adapter, args []string)
	priority  int
}

func spawnWave(w adapter.Adapter, args []string) {
	firstplayer, err := w.FirstPlayer()
	if err != nil {
		logrus.Errorf("unable to find player: %s", err)
		return
	}
	pPos := firstplayer.GetPosition()
	playerVector := vector.NewWithValues([]float64{pPos.X, pPos.Y})

	amount, _ := strconv.Atoi(args[1])
	hp, _ := strconv.Atoi(args[2])
	distance := 200

	radians_spread := (2.0 * math.Pi) / float64(amount)

	switch args[0] {
	case "radish":
		for i := 0; i < amount; i++ {
			x := math.Cos(radians_spread * float64(i))
			y := math.Sin(radians_spread * float64(i))
			spawnloc := vector.NewWithValues([]float64{x, y})
			spawnloc.Scale(float64(distance))
			spawnloc = vector.Add(spawnloc, playerVector)
			_, e := entity.MakeRandEntity(w, []string{
				"entities/radishred.entity.json",
				"entities/radishblue.entity.json",
				"entities/radishyellow.entity.json",
			}, spawnloc[0], spawnloc[1])
			radish, ok := e.(Enemy)
			if !ok {
				logrus.Error("some error with building radish entity")
				return
			}

			radish.GetHealth().Current = float64(hp)
		}
	default:
		return
	}
}

func spawnBoss(w adapter.Adapter, args []string) {
	// // hp, _ := strconv.Atoi(args[1])
	// distance, _ := strconv.Atoi(args[2])
	// loc_radian := rand.Float64() * (math.Pi * 2)

	// x := math.Cos(loc_radian)
	// y := math.Sin(loc_radian)

	// v := vector.NewWithValues([]float64{x, y})

	// v.Scale(float64(distance))

	// playerQuery := query.NewQuery(filter.Contains(tag.Player))

	// firstplayer, _ := playerQuery.FirstEntity(w)

	// pPos := component.GetPosition(firstplayer)
	// playerVector := vector.NewWithValues([]float64{pPos.X, pPos.Y})

	// v = vector.Add(playerVector, v)

	// entity.MakeEntity(w, "entities/cheezboss.entity.json", v[0], v[1])
}

func outputHurryUp(w adapter.Adapter, args []string) {
	logrus.Info("HURRY UP")
}

func outputDeath(w adapter.Adapter, args []string) {
	logrus.Info("Death")
}

type JobName string

const (
	SpawnWave JobName = "SpawnWave"
	SpawnBoss JobName = "SpawnBoss"
	HurryUp   JobName = "HurryUp"
	Death     JobName = "Death"
)

var JobTypes = make(map[JobName]Job, 1)

type SceneEvent struct {
	Time      uint32   `json:"time"`
	EventName JobName  `json:"event"`
	EventArgs []string `json:"args"`
}

func init() {
	JobTypes[HurryUp] = Job{
		json_name: string(HurryUp),
		Callback:  outputHurryUp,
		priority:  1,
	}

	JobTypes[Death] = Job{
		json_name: string(Death),
		Callback:  outputDeath,
		priority:  0,
	}

	JobTypes[SpawnBoss] = Job{
		json_name: string(SpawnBoss),
		Callback:  spawnBoss,
		priority:  0,
	}

	JobTypes[SpawnWave] = Job{
		json_name: string(SpawnWave),
		Callback:  spawnWave,
		priority:  0,
	}
}
