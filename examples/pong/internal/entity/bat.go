package entity

import "cheezewiz/examples/pong/internal/component"

type Bat struct {
	component.Rect
	component.Vel
	component.Pos
	component.IsPlayer
	component.IsBat
	// component.Controller
}
