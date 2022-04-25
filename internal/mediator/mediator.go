package mediator

import (
	"github.com/hajimehoshi/ebiten/v2"

	"cheesewiz/internal/chat"
	"cheesewiz/internal/input"
)

type Mediator struct {
	chat  *chat.Chat
	input input.Input
}

func New() Mediator {
	c := chat.New()
	return Mediator{
		chat: c,
		input: input.Input{
			Chat: c,
		},
	}
}
func (m *Mediator) Update() {
	m.input.Update()
	m.chat.Update()
}
func (m *Mediator) Render(dst *ebiten.Image) {
	m.chat.Render(dst)
}
