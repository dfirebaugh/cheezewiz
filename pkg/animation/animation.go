package animation

import (
	"cheezewiz/internal/filesystem"
	"image"

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

func MakeAnimation(path string, height int, width int) *Animation {
	a := &Animation{
		Path:        path,
		FrameWidth:  height,
		FrameHeight: width,
	}

	if len(a.Path) > 0 {
		a.Image = filesystem.GetPNG(path)
	}
	if a.Image == nil {
		logrus.Error("no source image for this animation")
		return nil
	}

	return a
}

func MakeDebugAnimation() *Animation {
	a := &Animation{
		Image:       ebiten.NewImage(10, 10),
		FrameWidth:  10,
		FrameHeight: 10,
	}

	a.Image.Fill(colornames.Yellowgreen)
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
