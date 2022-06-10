package gamemath

import (
	"cheezewiz/pkg/ecs"
	"math"

	"github.com/atedja/go-vector"
)

func GetHeading(a, b vector.Vector) float64 {
	r := vector.Unit(vector.Subtract(a, b))
	return math.Atan2(r[1], r[0])
}

func GetVector(a, b float64) vector.Vector {
	return []float64{a, b}
}

func GetDistance(a, b []float64) float64 {
	return math.Sqrt(math.Pow(a[0]-b[0], 2) + math.Pow(a[1]-b[1], 2))
}

func GetClosest(src vector.Vector, dest map[ecs.EntityHandle]vector.Vector) ecs.EntityHandle {
	var closestDistance float64 = 100000000
	var closestHandle ecs.EntityHandle

	for handle, v := range dest {
		distance := GetDistance(src, v)
		if distance < closestDistance {
			closestDistance = distance
			closestHandle = handle
		}
	}

	return closestHandle
}
