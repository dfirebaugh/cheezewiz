package input

type PlayerInput interface {
	IsUpPressed() bool
	IsDownPressed() bool
	IsLeftPressed() bool
	IsRightPressed() bool
	IsPrimaryAtkPressed() bool
	IsPrimaryAtkJustPressed() bool
}
