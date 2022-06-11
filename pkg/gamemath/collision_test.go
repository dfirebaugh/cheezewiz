package gamemath

import (
	"testing"
)

type test struct {
	a        Rect
	b        Rect
	expected bool
	message  string
}

func TestIsAxisAlignedCollision(t *testing.T) {
	tests := []test{
		{
			a:        Rect([]float64{1, 1, 1, 1}),
			b:        Rect([]float64{1, 1, 1, 1}),
			expected: true,
			message:  "1s should collide",
		},
		{
			a:        Rect([]float64{100, 100, 100, 100}),
			b:        Rect([]float64{1, 1, 1, 1}),
			expected: false,
			message:  "far away should not collide",
		},
		{
			a:        Rect([]float64{0, 0, 100, 100}),
			b:        Rect([]float64{100, 100, 10, 10}),
			expected: false,
			message:  "adjacent should not collide",
		},
		{
			a:        Rect([]float64{0, 0, 100, 100}),
			b:        Rect([]float64{98, 98, 100, 100}),
			expected: true,
			message:  "partial overlap should collide",
		},
	}

	for _, ts := range tests {
		if ts.a.IsAxisAlignedCollision(ts.b) != ts.expected {
			t.Errorf(ts.message)
			t.Fail()
		}
	}
}
