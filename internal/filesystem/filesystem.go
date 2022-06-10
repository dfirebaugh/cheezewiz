package filesystem

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/config"
	"embed"
	"image"
	_ "image/png"
	"sync"

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

var entityFS embed.FS = config.EntityFS

func GetEntity(path string) []byte {
	f, err := entityFS.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func GetAsset(path string) []byte {
	f, err := assetFS.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

var pngCache = map[string]*ebiten.Image{}
var pngCacheMut = sync.RWMutex{}

func GetPNG(path string) *ebiten.Image {
	pngCacheMut.Lock()
	defer pngCacheMut.Unlock()
	// check if we have the image in the cache
	if cachedIMG, ok := pngCache[path]; ok {
		return cachedIMG
	}

	img, _, err := image.Decode(bytes.NewReader(GetAsset(path)))
	if err != nil {
		logrus.Error(err)
		return nil
	}

	// save to cache
	pngCache[path] = ebiten.NewImageFromImage(img)
	return pngCache[path]
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
