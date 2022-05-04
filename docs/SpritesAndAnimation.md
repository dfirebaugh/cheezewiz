# Atlas
Ebiten automatically compiles our images to a global atlas of 4096x4096.  If we have more assets than 


# Animation
We are using [ganim8](https://github.com/yohamta/ganim8) to help easily make animations from spritesheets.

To make an animation, place a png of the sprite sheet in the `assets` directory.

The png should be in a grid layout (like the below image).
Animations are accessed by this function: `grid.GetFrames("1-5", 5)`
The first argument is the range in the column.  The second argument is the row.
So, it may make sense to have an action per row (e.g. the sequence for a jump might exist on one row and the sequence for an attack could exist on another)
> note that we should be able to trigger multiple rows in a single action

<p align="center">
  <img src="https://github.com/yohamta/ganim8/blob/master/examples/gif/example.gif?raw=true" />
</p>

<img src="https://github.com/yohamta/ganim8/blob/master/examples/assets/images/Character_Monster_Slime_Blue.png?raw=true" />

> images are from [ganim8](https://github.com/yohamta/ganim8) example

## Loading png into golang

golang allows us to [embed](https://pkg.go.dev/embed) into the binary.
e.g.

```golang
package assets


package icons

import (
	"bytes"
	_ "embed"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

// note that we don't actually call anything from the `embed`
//  package.  This package uses a special decorator comment that you can see below.  
//  e.g. //go:embed <name of file to embed>
//  The file has to exist in the same directory. I don't believe there is support for 
//  absolute paths.
//  Alternatively, we could choose to load the files at runtime.  If we don't want to embed the
//  binary.  This would require an entirely different implementation though.

//go:embed tabs.png
var tabsPng []byte

//go:embed icons.png
var iconsRaw []byte

//go:embed palette.png
var palletRaw []byte

var Tabs *ebiten.Image
var ToolBar *ebiten.Image
var Palette *ebiten.Image

func init() {
	Tabs, _ = loadPNG(tabsPng)
	ToolBar, _ = loadPNG(iconsRaw)
	Palette, _ = loadPNG(palletRaw)
}

func loadPNG(b []byte) (*ebiten.Image, error) {
	var err error
	imgDecoded, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return ebiten.NewImageFromImage(imgDecoded), nil
}
```
