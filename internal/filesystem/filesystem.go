package filesystem

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/config"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

// assetFS is relative to ./assets dir
var assetFS embed.FS = assets.AssetFS

// sceneFS is relative to ./assets/scenes dir
// this will contain map data
var sceneFS embed.FS = assets.SceneFS

// levelFS is relative to ./config/levels dir
// this will contain metadata about a level e.g. what spawns trigger at what time
var levelFS embed.FS = config.LevelFS

func GetAsset(path string) []byte {
	f, err := assetFS.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func GetPNG(path string) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(GetAsset(path)))
	if err != nil {
		logrus.Error(err)
	}
	return ebiten.NewImageFromImage(img)
}

func GetLevel(path string) []byte {
	f, err := levelFS.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func GetScene(path string) []byte {
	f, err := sceneFS.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}
