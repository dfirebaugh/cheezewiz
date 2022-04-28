package console

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type consoleService interface {
	Submit(input string)
	GetResponse() string
}
type Console struct {
	runes    []rune
	text     string
	response string
	counter  int
	isOpen   bool
	service  consoleService
}

func New(service consoleService) *Console {
	c := &Console{
		text:    "",
		counter: 0,
		isOpen:  false,
		service: service,
	}
	return c
}

func (c *Console) Update() error {
	c.handleEnter()
	c.handleSlash()
	c.handleResponse()

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

func (c *Console) Render(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, c.response)

	if !c.isOpen {
		return
	}
	// Blink the cursor.
	t := c.text
	if c.counter%60 < 30 {
		t += "_"
	}
	ebitenutil.DebugPrintAt(screen, t, 0, screen.Bounds().Dy()-20)
}

func (c *Console) IsOpen() bool {
	return c.isOpen
}

func (c *Console) SetOpen(open bool) {
	c.isOpen = open
}

func (c *Console) ToggleOpen() {
	c.isOpen = !c.isOpen
}

func (c *Console) isWithinCharLimit() bool {
	return len(c.text) > 40
}

func (c *Console) handleInput() {
	// Add runes that are input by the user by AppendInputChars.
	// Note that AppendInputChars result changes every frame, so you need to call this
	// every frame.
	c.runes = ebiten.AppendInputChars(c.runes[:0])
	c.text += string(c.runes)

	c.counter++
}

func (c *Console) handleEnter() {
	// If the enter key is pressed, submit
	if !(repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter)) {
		return
	}
	if len(c.text) == 0 {
		return
	}
	c.Submit()
}

func (c *Console) handleSlash() {
	// If the / key is pressed, submit
	if !repeatingKeyPressed(ebiten.KeySlash) {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		return
	}

	if c.IsOpen() {
		return
	}

	if len(c.text) == 0 {
		return
	}

	if c.text == "/" {
		return
	}

	c.text = "/"
	c.Submit()
}

func (c *Console) handleBackSpace() {
	// If the backspace key is pressed, remove one character.
	if !repeatingKeyPressed(ebiten.KeyBackspace) {
		return
	}
	if len(c.text) < 1 {
		return
	}
	c.text = c.text[:len(c.text)-1]
}

func (c *Console) Submit() {
	c.service.Submit(c.text)
	c.clear()
	c.handleResponse()
}

func (c *Console) clear() {
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

func (c *Console) handleResponse() {
	r := c.service.GetResponse()
	if len(r) == 0 {
		return
	}
	c.response = r
}
