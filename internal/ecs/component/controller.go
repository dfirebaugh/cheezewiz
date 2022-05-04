package component

type PlayerInput interface {
	IsUpPressed() bool
	IsDownPressed() bool
	IsLeftPressed() bool
	IsRightPressed() bool
}

type Controller struct {
	Controller PlayerInput
}
