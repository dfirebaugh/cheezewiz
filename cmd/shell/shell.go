package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"cheezewiz/internal/services/chatservice"
	"cheezewiz/internal/services/console"

	"github.com/sirupsen/logrus"
)

type consoleService interface {
	Submit(string)
	GetResponse() string
	GetCommands() map[string]func() string
}

type shell struct {
	text           string
	lastResponse   string
	scanner        *bufio.Reader
	service        consoleService
	responseBuffer *strings.Reader
}

func (shell) printRepl() {
	fmt.Print("cheeze> ")
}
func (sh *shell) shouldContinue() bool {
	return !strings.EqualFold("exit", sh.text)
}

func (sh *shell) newReader() {
	sh.responseBuffer = strings.NewReader(sh.lastResponse)
	sh.scanner = bufio.NewReader(os.Stdin)
}

func (sh *shell) iter() {
	sh.printRepl()
	t, _ := sh.scanner.ReadString('\n')
	sh.text = strings.TrimSpace(t)
}

func (sh *shell) cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (sh *shell) shouldClearLastReponse() bool {
	if _, exists := sh.service.GetCommands()[strings.Split(sh.text, " ")[0]]; exists {
		return false
	}

	return true
}

func (sh *shell) hasNewResponse(count int) bool {
	if count == 100 {
		logrus.Warnf("didn't have a a response after %d attempst", count)
		return true
	}
	currentResponse := sh.service.GetResponse()
	if len(currentResponse) > 0 && sh.lastResponse != currentResponse {
		sh.lastResponse = currentResponse
		println(sh.lastResponse)
		return true
	}
	return false
}

func (sh *shell) submit() bool {
	if sh.text == "" {
		return false
	}

	sh.service.Submit(sh.text)
	sh.text = ""
	return true
}

func (sh *shell) run() {
	for ; sh.shouldContinue(); sh.iter() {
		if sh.text == "clear" {
			sh.cls()
			continue
		}
		if sh.shouldClearLastReponse() {
			sh.lastResponse = ""
		}
		sh.hasNewResponse(0)

		lookForResponse := sh.submit()

		if !lookForResponse {
			continue
		}

		count := 0
		for {
			if sh.hasNewResponse(count) {
				break
			}
			count++
		}
	}
	println("Bye!")
}

func main() {
	logrus.SetLevel(logrus.WarnLevel)
	sh := shell{
		service: console.New(chatservice.NewClient()),
	}
	sh.newReader()

	sh.run()
}
