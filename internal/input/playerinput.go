package input

type PlayerInput interface {
	IsUpPressed() bool
	IsDownPressed() bool
	IsLeftPressed() bool
	IsLeftJustPressed() bool
	IsRightPressed() bool
	IsRightJustPressed() bool
	IsPrimaryAtkPressed() bool
	IsPrimaryAtkJustPressed() bool
}
