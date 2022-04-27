package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type chat interface {
	ToggleOpen()
	Open()
	IsOpen() bool
}

type Input struct {
	Chat chat
}

func (Input) isEnterJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}

func (Input) isSlashJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySlash)
}

func (i Input) Update() {
	if i.isSlashJustPressed() {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			return
		}

		if i.Chat.IsOpen() {
			return
		}

		i.Chat.Open()
	}
	if i.isEnterJustPressed() {
		i.Chat.ToggleOpen()
	}
}
