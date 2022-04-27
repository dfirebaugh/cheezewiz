package console

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
)

type Console struct {
	runes       []rune
	text        string
	response    string
	counter     int
	isOpen      bool
	chatClient  ChatClient
	username    string
	destination string
}

type ChatClient interface {
	PublishMessage(string, string, string) string
	SubscribeMessage(source string, destination string) error
	GetResponse() string
}

func New(chatClient ChatClient) *Console {
	// err := chatClient.SubscribeMessage("default", "world")
	// if err != nil {
	// 	logrus.Errorf("error in console instantiation: %s", err)
	// }
	cc := chatClient
	c := &Console{
		text:        "",
		counter:     0,
		isOpen:      false,
		chatClient:  cc,
		username:    "default",
		destination: "world",
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
	ebitenutil.DebugPrintAt(screen, t, 0, 230)
	// ebitenutil.DebugPrintAt(screen, c.chatClient.GetResponse(), 10, 230)
}

func (c *Console) Open() {
	c.isOpen = true
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
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(c.text) >= 1 {
			c.text = c.text[:len(c.text)-1]
		}
	}
}

func (c *Console) Submit() {
	logrus.Info(c.text)

	if strings.Split(c.text, "")[0] == "/" {
		logrus.Info("evaluatinng a command")
		c.evaluateCommand()
		c.clear()
		return
	}

	c.sendMessage(c.text)
	c.clear()
}

func (c *Console) clear() {
	c.text = ""
	c.counter = 0
}

func (c *Console) evaluateCommand() {
	split := strings.Split(c.text, " ")

	if split[0] == "/who" {
		c.response = c.username
		return
	}
	if split[0] == "/u" {
		c.setUserName(split[1])
		return
	}
	if split[0] == "/r" {
		c.setDestination(split[1])
		return
	}

}

func (c *Console) setUserName(username string) {
	if username == "" {
		return
	}

	c.username = username
	c.response = fmt.Sprintf("user set [%s]", c.username)
	logrus.Info("setting username to: ", c.username)
	err := c.chatClient.SubscribeMessage(c.username, c.username)
	if err != nil {
		logrus.Errorf("error in console instantiation: %s", err)
	}
	err = c.chatClient.SubscribeMessage(c.username, "world")
	if err != nil {
		logrus.Errorf("error in console instantiation: %s", err)
	}
}

func (c *Console) setDestination(destination string) {
	if destination == "" {
		return
	}

	c.destination = destination
	c.response = fmt.Sprintf("user set [%s]", c.destination)
	logrus.Info("setting destination to: ", c.destination)
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

func (c *Console) sendMessage(content string) {
	c.chatClient.PublishMessage(c.username, c.destination, content)
}

func (c *Console) handleResponse() {
	r := c.chatClient.GetResponse()
	if len(r) == 0 {
		return
	}
	c.response = r
}
