package gamemath

import (
	"fmt"
	"math"

	"github.com/atedja/go-vector"
)

// Vector a float64 slice with 2 elements []float64{x, y}
type Vector []float64

func MakeVector(x, y float64) Vector {
	return Vector([]float64{x, y})
}

func (v Vector) ToString() string {
	return fmt.Sprintf("%d, %d", int(v[0]), int(v[1]))
}

// origin (0,0) is top left of screen
func (v Vector) ToTileCoord(tileSize float64) Vector {
	return Vector([]float64{v[0] / tileSize, v[1] / tileSize})
}

func (v Vector) IsOffScreen(screenHeight float64, screenWidth float64) bool {
	return v[0] < 0 || v[1] < 0 || v[0] > screenWidth || v[1] > screenHeight
}

func (v Vector) ToLocation() (int, int) {
	return int(v[0]), int(v[1])
}

func (v Vector) Offset(o Vector) Vector {
	return Vector([]float64{v[0] - o[0], v[1] - o[1]})
}

func (v Vector) GetDistance(b Vector) float64 {
	return math.Sqrt(math.Pow(v[0]-b[0], 2) + math.Pow(v[1]-b[1], 2))
}

// returns radian toward vector
func (v Vector) GetHeading(target Vector) float64 {
	r := vector.Unit(vector.Subtract(vector.Vector(v), vector.Vector(target)))
	return math.Atan2(r[1], r[0])
}

func (v Vector) Scale(value float64) {
	vector.Vector(v).Scale(value)
}

func (v Vector) Add(o Vector) Vector {
	return Vector(vector.Add(vector.Vector(v), vector.Vector(o)))
}
