package gamemath

import (
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

func GetClosest[T comparable](src vector.Vector, dest map[T]vector.Vector) T {
	var closestDistance float64 = 100000000
	var closestHandle T

	for handle, v := range dest {
		distance := GetDistance(src, v)
		if distance < closestDistance {
			closestDistance = distance
			closestHandle = handle
		}
	}

	return closestHandle
}

// Rect a float64 slice with 4 elements []float64{x, y, width, height}
type Rect []float64

func (r Rect) IsAxisAlignedCollision(other Rect) bool {
	ax := r[0]
	ay := r[1]
	aw := r[2]
	ah := r[3]

	bx := other[0]
	by := other[1]
	bw := other[2]
	bh := other[3]

	return ax < bx+bw &&
		ax+aw > bx &&
		ay < by+bh &&
		ah+ay > by
}
