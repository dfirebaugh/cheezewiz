package component

import "github.com/yohamta/donburi"

type RadiusData struct {
	Value float64
}

var Radius = donburi.NewComponentType(RadiusData{})

func GetRadius(entry *donburi.Entry) *RadiusData {
	return (*RadiusData)(entry.Component(Radius))
}

type RectData struct {
	Height float64
	Width  float64
}

var Rect = donburi.NewComponentType(RectData{})

func GetRect(entry *donburi.Entry) *RectData {
	return (*RectData)(entry.Component(Rect))
}
