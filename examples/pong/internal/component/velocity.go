package component

// Position for any entity, if it needs
type Pos struct {
	X, Y float64 // Just a 2D point
}

// Velocity for any entity, if it needs
type Vel struct {
	L, M float64 // Also, 2D point
}

// Radius for any entity, if it needs
type Rad struct {
	Value float64 // Width value
}
