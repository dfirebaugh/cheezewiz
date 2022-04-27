package chatservice

import (
	"cheezewiz/internal/models/message"
	"cheezewiz/pkg/broker"
	"context"
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/sirupsen/logrus"
)

const Port = 10578

type streamHandler struct {
	wg     *sync.WaitGroup
	stream message.MessageHandler_SubscribeMessageServer
}

type chatServer struct {
	streams    map[string][]streamHandler
	chatBroker *broker.Broker
}

func NewServer() *chatServer {
	b := broker.New()
	go b.Start()
	return &chatServer{
		streams:    make(map[string][]streamHandler),
		chatBroker: b,
	}
}

func (c *chatServer) PublishMessage(ctx context.Context, request *message.MessageRequest) (*flatbuffers.Builder, error) {
	logrus.Debugf("publishing message: [%s->%s]: %s", string(request.Source()), string(request.Destination()), string(request.Body()))

	for _, s := range c.streams[string(request.Destination())] {
		s.wg.Add(1)
	}

	go c.chatBroker.Publish(brokerMessage{
		destination: string(request.Destination()),
		source:      string(request.Source()),
		message:     string(request.Body()),
	})
	return buildMessageFromRequest(request), nil
}

func (c *chatServer) registerListener(destination string, stream message.MessageHandler_SubscribeMessageServer) {
	logrus.Infof("registering listener for: %s", destination)
	_, exists := c.streams[destination]
	if !exists {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.streams[destination] = []streamHandler{{
			wg:     wg,
			stream: stream,
		}}
		return
	}

	sh := streamHandler{
		wg:     &sync.WaitGroup{},
		stream: stream,
	}
	sh.wg.Add(1)

	c.streams[destination] = append(c.streams[destination], sh)
}

func (c *chatServer) SubscribeMessage(request *message.MessageRequest, stream message.MessageHandler_SubscribeMessageServer) error {
	logrus.Debugf("rcvd a sub to: [%s] from [%s]", string(request.Destination()), string(request.Source()))

	c.registerListener(string(request.Destination()), stream)
	go c.listen()

	for _, s := range c.streams[string(request.Destination())] {
		s.wg.Wait()
	}

	return nil
}

func (c *chatServer) listen() {
	logrus.Debug("chat service running")

	msg := c.chatBroker.Subscribe()

	for {
		m := <-msg
		// defer mbx.wg.Done()
		logrus.Infof("[mailbox] sending: [%s->%s]: %s", m.GetRequestor(), m.GetTopic(), m.GetPayload())
		builtMsg := buildMessage(m.GetRequestor(), m.GetTopic(), m.GetPayload().(string))
		for _, s := range c.streams[m.GetTopic()] {
			s.stream.Send(builtMsg)
		}
	}
}
