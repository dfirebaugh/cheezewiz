package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type chat interface {
	ToggleOpen()
	SetOpen(bool)
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

		i.Chat.SetOpen(true)
	}
	if i.isEnterJustPressed() {
		i.Chat.ToggleOpen()
	}
}
