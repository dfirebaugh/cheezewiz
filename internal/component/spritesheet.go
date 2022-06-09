package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	IMG  *ebiten.Image `json:"img"`
	Path string        `json:"path"`
}
