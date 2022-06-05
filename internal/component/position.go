package component

import (
	"github.com/yohamta/donburi"
)

type PositionData struct {
	X  float64
	Y  float64
	CX float64
	CY float64
}

var Position = donburi.NewComponentType(PositionData{})

func GetPosition(entry *donburi.Entry) *PositionData {
	return (*PositionData)(entry.Component(Position))
}

func (p *PositionData) Set(x float64, y float64, cx float64, cy float64) {
	p.X = x
	p.Y = y
	p.CX = cx
	p.CY = cy
}

func (p *PositionData) Update(x float64, y float64) {
	p.X = x
	p.Y = y
}
