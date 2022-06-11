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

type GameFileSystem struct {
	asset       embed.FS
	scene       embed.FS
	level       embed.FS
	entity      embed.FS
	pngCache    map[string]*ebiten.Image
	pngCacheMut sync.RWMutex
}

var Game = GameFileSystem{
	// assetFS is relative to ./assets dir
	asset: assets.AssetFS,
	// sceneFS is relative to ./assets/scenes dir
	// this will contain map data
	scene: assets.SceneFS,
	// levelFS is relative to ./config/levels dir
	// this will contain metadata about a level e.g. what spawns trigger at what time
	level:       config.LevelFS,
	entity:      config.EntityFS,
	pngCache:    map[string]*ebiten.Image{},
	pngCacheMut: sync.RWMutex{},
}

func (fs *GameFileSystem) GetEntity(path string) []byte {
	f, err := fs.entity.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func (fs *GameFileSystem) GetAsset(path string) []byte {
	f, err := fs.asset.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func (fs *GameFileSystem) GetPNG(path string) *ebiten.Image {
	fs.pngCacheMut.Lock()
	defer fs.pngCacheMut.Unlock()
	// check if we have the image in the cache
	if cachedIMG, ok := fs.pngCache[path]; ok {
		return cachedIMG
	}

	img, _, err := image.Decode(bytes.NewReader(fs.GetAsset(path)))
	if err != nil {
		logrus.Error(err)
		return nil
	}

	// save to cache
	fs.pngCache[path] = ebiten.NewImageFromImage(img)
	return fs.pngCache[path]
}

func (fs *GameFileSystem) GetLevel(path string) []byte {
	f, err := fs.level.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func (fs *GameFileSystem) GetScene(path string) []byte {
	f, err := fs.scene.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}
