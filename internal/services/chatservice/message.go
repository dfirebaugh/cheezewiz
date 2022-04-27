package chatservice

import (
	"cheezewiz/internal/models/message"
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
)

type brokerMessage struct {
	destination string
	source      string
	message     string
}

func (m brokerMessage) GetPayload() interface{} {
	return m.message
}
func (m brokerMessage) GetRequestor() string {
	return m.source
}
func (m brokerMessage) GetTopic() string {
	return m.destination
}
func (m brokerMessage) Hash() string {
	return m.String()
}
func (m brokerMessage) String() string {
	return fmt.Sprintf("%s|%s|%s", m.source, m.destination, m.message)
}

func buildMessage(source string, destination string, content string) *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)

	s := b.CreateString(source)
	d := b.CreateString(destination)
	m := b.CreateString(content)

	message.MessageRequestStart(b)
	message.MessageRequestAddDestination(b, d)
	message.MessageRequestAddSource(b, s)
	message.MessageRequestAddBody(b, m)

	b.Finish(message.MessageRequestEnd(b))

	return b
}

func buildMessageFromRequest(request *message.MessageRequest) *flatbuffers.Builder {
	bodyRaw := request.Body()
	sourceRaw := request.Source()
	destinationRaw := request.Destination()

	var m string
	var s string
	var d string
	if bodyRaw == nil || sourceRaw == nil || destinationRaw == nil {
		m = "Unknown"
		s = "Unknown"
		d = "Unknown"
	} else {
		m = string(bodyRaw)
		s = string(sourceRaw)
		d = string(destinationRaw)
	}

	b := flatbuffers.NewBuilder(0)
	messageContent := b.CreateString(m)
	messageDestination := b.CreateString(d)
	messageSource := b.CreateString(s)
	message.MessageRequestStart(b)
	message.MessageRequestAddBody(b, messageContent)
	message.MessageRequestAddSource(b, messageSource)

	message.MessageRequestAddDestination(b, messageDestination)
	b.Finish(message.MessageRequestEnd(b))

	return b
}
