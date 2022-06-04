package event

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"math"
	"math/rand"
	"strconv"

	"github.com/atedja/go-vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Job struct {
	json_name string
	callback  func(w donburi.World, args []string)
	priority  int
}

func spawnWave(w donburi.World, args []string) {
	playerQuery := query.NewQuery(filter.Contains(entity.PlayerTag))

	firstplayer, _ := playerQuery.FirstEntity(w)

	pPos := component.GetPosition(firstplayer)
	playerVector := vector.NewWithValues([]float64{pPos.X, pPos.Y})

	amount, _ := strconv.Atoi(args[1])
	hp, _ := strconv.Atoi(args[2])
	distance := 200

	radians_spread := (2.0 * math.Pi) / float64(amount)

	switch {
	case "radish" == args[0]:
		for i := 0; i < amount; i++ {
			x := math.Cos(radians_spread * float64(i))
			y := math.Sin(radians_spread * float64(i))
			spawnloc := vector.NewWithValues([]float64{x, y})
			spawnloc.Scale(float64(distance))
			spawnloc = vector.Add(spawnloc, playerVector)
			e := entity.MakeEnemy(w, spawnloc[0], spawnloc[1])
			component.GetHealth(e).HP = float64(hp)
		}

	default:
		return
	}
}
func spawnBoss(w donburi.World, args []string) {

	hp, _ := strconv.Atoi(args[1])
	distance, _ := strconv.Atoi(args[2])
	loc_radian := rand.Float64() * (math.Pi * 2)

	x := math.Cos(loc_radian)
	y := math.Sin(loc_radian)

	v := vector.NewWithValues([]float64{x, y})

	v.Scale(float64(distance))

	playerQuery := query.NewQuery(filter.Contains(entity.PlayerTag))

	firstplayer, _ := playerQuery.FirstEntity(w)

	pPos := component.GetPosition(firstplayer)
	playerVector := vector.NewWithValues([]float64{pPos.X, pPos.Y})

	v = vector.Add(playerVector, v)

	e := entity.MakeEnemy(w, v[0], v[1])
	component.GetHealth(e).HP = float64(hp)
	//args[]
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
