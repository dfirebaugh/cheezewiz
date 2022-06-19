package gamemath

import "fmt"

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
