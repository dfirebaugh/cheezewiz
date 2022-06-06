package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteSheetData struct {
	IMG  *ebiten.Image `json:"img"`
	Path string        `json:"path"`
}

var SpriteSheet = donburi.NewComponentType(SpriteSheetData{})

func GetSpriteSheet(entry *donburi.Entry) *SpriteSheetData {
	return (*SpriteSheetData)(entry.Component(SpriteSheet))
}

var ImageOptions = donburi.NewComponentType(ebiten.DrawImageOptions{})

func GetImageOptions(entry *donburi.Entry) *ebiten.DrawImageOptions {
	return (*ebiten.DrawImageOptions)(entry.Component(ImageOptions))
}
