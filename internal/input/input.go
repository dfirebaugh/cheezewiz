package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type chat interface {
	ToggleOpen()
}

type Input struct {
	Chat chat
}

func (Input) isEnterJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}

func (i Input) Update() {
	if i.isEnterJustPressed() {
		i.Chat.ToggleOpen()
	}
}
