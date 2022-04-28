package console

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type ChatClient interface {
	PublishMessage(string, string, string) string
	SubscribeMessage(source string, destination string) error
	GetResponse() string
}

type input []string

type console struct {
	input       input
	response    string
	username    string
	destination string
	chatClient  ChatClient
}

func (ci input) isCommand() bool {
	first := ci[0]
	return string(first[0]) == "/"
}
func (ci input) hasArgs() bool {
	return len(ci) > 1
}

func New(chatClient ChatClient) *console {
	return &console{
		chatClient:  chatClient,
		username:    "default",
		destination: "world",
	}
}

func (c *console) Submit(input string) {
	c.input = strings.Split(input, " ")
	logrus.Info(c.input)

	if c.input.isCommand() {
		logrus.Info("evaluating a command")
		c.evaluateCommand()
		c.clear()
		return
	}

	c.SendMessage(strings.Join(c.input, " "))
	c.clear()
}

func (c *console) GetResponse() string {
	c.handleResponse()
	return c.response
}

func (c *console) ClearResponse() {
	c.response = ""
}
func (c *console) GetCommands() map[string]func() string {
	return map[string]func() string{
		"/who": c.getUsername,
		"/u":   c.setUserName,
	}
}

func (c *console) clear() {
	c.input = []string{}
}

func (c *console) evaluateCommand() {
	if cmd, exists := c.GetCommands()[c.input[0]]; exists {
		logrus.Infof("running command: %s", c.input[0])
		c.response = cmd()
		return
	}
	logrus.Warnf("command doesn't exist: %s", c.input[0])
}

func (c *console) setUserName() string {
	if !c.input.hasArgs() {
		logrus.Warn("no username set")
		return ""
	}
	if c.input[1] == "" {
		return ""
	}

	c.username = c.input[1]
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

	return fmt.Sprintf("set current user to [%s]", c.username)
}

func (c *console) getUsername() string {
	return fmt.Sprintf("current user: [%s]", c.username)
}

func (c *console) SendMessage(msg string) {
	c.chatClient.PublishMessage(c.username, c.destination, msg)
}

func (c *console) handleResponse() {
	r := c.chatClient.GetResponse()
	if len(r) == 0 {
		return
	}
	c.response = r
}
