package component

type InputController interface {
	IsUpPressed() bool
	IsDownPressed() bool
	IsLeftPressed() bool
	IsRightPressed() bool
}

type Controller struct {
	Controller InputController
}
