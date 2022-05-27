package entity

import (
	"cheezewiz/examples/pong/internal/component"
)

// Your game object
type Ball struct {
	component.Pos // Ball position
	component.Vel // Ball velocity
	component.Rad // Ball radius
}
