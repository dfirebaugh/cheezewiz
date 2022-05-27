package entity

import "cheezewiz/examples/pong/internal/component"

type Border struct {
	component.Rect
	component.Pos
	component.IsBat
}
