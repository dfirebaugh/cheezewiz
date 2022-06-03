package entity

import (
	"bytes"
	"cheezewiz/examples/choppa/assets"
	"cheezewiz/examples/choppa/internal/component"
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var FighterTag = donburi.NewTag()

func MakeFighter(w donburi.World) {
	screenWidth, screenHeight := ebiten.WindowSize()

	p := w.Create(FighterTag, component.Position, component.Velocity, component.SpriteSheet, component.IsAlive)
	entry := w.Entry(p)
	position := (*component.PositionData)(entry.Component(component.Position))
	velocity := (*component.VelocityData)(entry.Component(component.Velocity))
	spriteSheet := (*component.SpriteSheetData)(entry.Component(component.SpriteSheet))
	alive := (*component.AliveData)(entry.Component(component.IsAlive))

	alive.IsAlive = true

	*velocity = component.VelocityData{
		L: -2,
		M: 0,
	}

	*position = component.PositionData{
		X: float64(screenWidth / 2),
		Y: float64(rand.Intn(screenHeight - 0)),
	}

	imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.FighterRaw))
	spriteSheet.IMG = ebiten.NewImageFromImage(imgDecoded)
}
