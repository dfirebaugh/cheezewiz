package system

import (
	"cheezewiz/pkg/ecs"
	"testing"
)

type test struct {
	a        []float64
	b        []float64
	expected bool
	message  string
}

func TestIsCollide(t *testing.T) {
	c := NewCollision(ecs.NewWorld())
	tests := []test{
		{
			a:        []float64{1, 1, 1, 1},
			b:        []float64{1, 1, 1, 1},
			expected: true,
			message:  "1s should collide",
		},
		{
			a:        []float64{100, 100, 100, 100},
			b:        []float64{1, 1, 1, 1},
			expected: false,
			message:  "far away should not collide",
		},
		{
			a:        []float64{0, 0, 100, 100},
			b:        []float64{100, 100, 10, 10},
			expected: false,
			message:  "adjacent should not collide",
		},
		{
			a:        []float64{0, 0, 100, 100},
			b:        []float64{98, 98, 100, 100},
			expected: true,
			message:  "partial overlap should collide",
		},
	}

	for _, ts := range tests {
		if c.IsCollide(ts.a, ts.b) != ts.expected {
			t.Errorf(ts.message)
			t.Fail()
		}
	}

}
