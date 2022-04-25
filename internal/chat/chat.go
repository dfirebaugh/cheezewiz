package chat

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
)

type Chat struct {
	runes   []rune
	text    string
	counter int
	isOpen  bool
}

func (c *Chat) Update() error {
	c.handleEnter()

	if !c.isOpen {
		return nil
	}

	c.handleBackSpace()

	// only allow a certain amount of characters
	if c.isWithinCharLimit() {
		return nil
	}

	c.handleInput()

	return nil
}

func (c *Chat) Render(screen *ebiten.Image) {
	if !c.isOpen {
		return
	}
	// Blink the cursor.
	t := c.text
	if c.counter%60 < 30 {
		t += "_"
	}
	ebitenutil.DebugPrintAt(screen, t, 0, 230)
}

func New() *Chat {
	return &Chat{
		text:    "",
		counter: 0,
		isOpen:  false,
	}
}

func (c *Chat) ToggleOpen() {
	c.isOpen = !c.isOpen
}

func (c *Chat) isWithinCharLimit() bool {
	return len(c.text) > 40
}

func (c *Chat) handleInput() {
	// Add runes that are input by the user by AppendInputChars.
	// Note that AppendInputChars result changes every frame, so you need to call this
	// every frame.
	c.runes = ebiten.AppendInputChars(c.runes[:0])
	c.text += string(c.runes)

	c.counter++
}

func (c *Chat) handleEnter() {
	// If the enter key is pressed, submit
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		if len(c.text) == 0 {
			return
		}
		c.Submit()
	}
}

func (c *Chat) handleBackSpace() {
	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(c.text) >= 1 {
			c.text = c.text[:len(c.text)-1]
		}
	}
}
func (c *Chat) Submit() {
	logrus.Info(c.text)
	c.text = ""
	c.counter = 0
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
