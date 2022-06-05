package component

import (
	"bytes"
	"cheezewiz/assets"
	"cheezewiz/internal/constant"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type Action struct {
	Sprite    *ganim8.Sprite
	Animation *ganim8.Animation
}

type AnimationData struct {
	Animations     map[ActorStateType]Action
	PrevUpdateTime time.Time
}

var Animation = donburi.NewComponentType(AnimationData{
	Animations: map[ActorStateType]Action{},
})

func GetAnimation(entry *donburi.Entry) *AnimationData {
	return (*AnimationData)(entry.Component(Animation))
}

func (a AnimationData) Get(state ActorStateType) Action {
	if _, ok := a.Animations[state]; !ok {
		imgDecoded, _, _ := image.Decode(bytes.NewReader(assets.CheezeWizRaw))
		grid := ganim8.NewGrid(int(constant.SpriteSize), int(constant.SpriteSize), imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

		return Action{
			Sprite:    ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1", 1)),
			Animation: ganim8.NewAnimation(ganim8.NewSprite(ebiten.NewImageFromImage(imgDecoded), grid.GetFrames("1", 1)), 100*time.Millisecond, ganim8.Nop),
		}
	}
	return a.Animations[state]
}
