package animation

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/colornames"
)

type Animation struct {
	Path        string
	Image       *ebiten.Image
	FrameWidth  int
	FrameHeight int
	FrameCount  int
	iter        int
	OX          int
	OY          int
}

type fs interface {
	GetPNG(path string) *ebiten.Image
}

func MakeAnimation(path string, height int, width int, fs fs) *Animation {
	a := &Animation{
		Path:        path,
		FrameWidth:  height,
		FrameHeight: width,
	}

	if len(a.Path) > 0 {
		a.Image = fs.GetPNG(path)
	}
	if a.Image == nil {
		logrus.Error("no source image for this animation")
		return nil
	}

	return a
}

// var debugImg =
var possibleColors = []color.Color{
	colornames.Violet,
	colornames.Tomato,
	colornames.Orange,
	colornames.Red,
	colornames.Purple,
	colornames.Green,
	colornames.Yellow,
}

func MakeDebugAnimation() *Animation {
	a := &Animation{
		Image:       ebiten.NewImage(10, 10),
		FrameWidth:  10,
		FrameHeight: 10,
	}

	a.Image.Fill(possibleColors[rand.Intn(len(possibleColors))])
	return a
}

func (a *Animation) GetFrame() *ebiten.Image {
	if a.Image == nil {
		logrus.Errorf("error getting Image")
		return nil
	}
	sx, sy := a.OX+a.getFrameIndex()*a.FrameWidth, a.OY
	return a.Image.SubImage(image.Rect(sx, sy, sx+a.FrameWidth, sy+a.FrameHeight)).(*ebiten.Image)
}

func (a *Animation) getFrameIndex() int {
	return (a.iter / 5) % a.getFrameCount()
}

func (a *Animation) getFrameCount() int {
	return a.Image.Bounds().Max.X / a.FrameWidth
}

func (a *Animation) IterFrame() {
	a.iter++
}
