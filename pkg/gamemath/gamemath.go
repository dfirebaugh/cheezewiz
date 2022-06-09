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

func GetClosest(src vector.Vector, dest map[int]vector.Vector) int {
	var closestDistance float64 = 100000000
	var closestHandle int

	for handle, v := range dest {
		distance := GetDistance(src, v)
		if distance < closestDistance {
			closestDistance = distance
			closestHandle = handle
		}
	}

	return closestHandle
}

// var closestEntry *donburi.Entry
// 			var closestDistance float64 = 100000000

// 			destinationQuery.EachEntity(w, func(pentry *donburi.Entry) {
// 				enemyPosition := component.GetPosition(pentry)
// 				if closestEntry == nil && w.Valid(pentry.Entity()) {
// 					closestEntry = pentry
// 				} else {
// 					distance := gamemath.GetDistance([]float64{position.X, position.Y}, []float64{enemyPosition.X, enemyPosition.Y})
// 					if distance < closestDistance && w.Valid(pentry.Entity()) {
// 						closestDistance = distance
// 						closestEntry = pentry
// 					}
// 				}
// 			})
