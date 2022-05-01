package net

import (
	"cheezewiz/internal/net/models/message"

	flatbuffers "github.com/google/flatbuffers/go"
)

func BuildMessage(source string, destination string, content string) *flatbuffers.Builder {
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

func BuildMessageFromRequest(request *message.MessageRequest) *flatbuffers.Builder {
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
