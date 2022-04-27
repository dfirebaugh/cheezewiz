package mediator

import (
	"github.com/hajimehoshi/ebiten/v2"

	"cheezewiz/internal/console"
	"cheezewiz/internal/input"
	"cheezewiz/internal/services/chatservice"
)

type Mediator struct {
	console *console.Console
	input   input.Input
}

func New() Mediator {
	c := console.New(chatservice.NewClient())
	return Mediator{
		console: c,
		input: input.Input{
			Chat: c,
		},
	}
}

func (m *Mediator) Update() {
	m.input.Update()
	m.console.Update()
}

func (m *Mediator) Render(dst *ebiten.Image) {
	m.console.Render(dst)
}
