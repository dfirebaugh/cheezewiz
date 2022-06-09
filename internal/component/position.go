package component

type Position struct {
	X  float64
	Y  float64
	CX float64
	CY float64
}

func (p *Position) Set(x float64, y float64, cx float64, cy float64) {
	p.X = x
	p.Y = y
	p.CX = cx
	p.CY = cy
}

func (p *Position) Update(x float64, y float64) {
	p.X = x
	p.Y = y
}
