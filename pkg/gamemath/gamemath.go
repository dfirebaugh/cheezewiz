package gamemath

import (
	"cheezewiz/pkg/ecs"
	"fmt"
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

// Vector a float64 slice with 2 elements []float64{x, y}
type Vector []float64

// Rect a float64 slice with 4 elements []float64{x, y, width, height}
type Rect []float64

func (v Vector) ToString() string {
	return fmt.Sprintf("%d, %d", int(v[0]), int(v[1]))
}

// origin (0,0) is top left of screen
func (v Vector) ToTileCoord(tileSize float64) Vector {
	return Vector([]float64{v[0] / tileSize, v[1] / tileSize})
}

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
