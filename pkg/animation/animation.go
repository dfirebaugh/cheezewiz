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
	FrameNum    int
	count       int
	OX          int
	OY          int
}

func MakeAnimation(path string, height int, width int) *Animation {
	a := &Animation{
		Path:        path,
		FrameWidth:  height,
		FrameHeight: width,
		OX:          0,
		OY:          0,
		count:       0,
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
	a.FrameNum = a.Image.Bounds().Max.X / a.FrameWidth
	i := (a.count / 5) % a.FrameNum
	sx, sy := a.OX+i*a.FrameWidth, a.OY
	return a.Image.SubImage(image.Rect(sx, sy, sx+a.FrameWidth, sy+a.FrameHeight)).(*ebiten.Image)
}

func (a *Animation) NextFrame() {
	a.count++
}
