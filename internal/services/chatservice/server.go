package chatservice

import (
	"cheezewiz/internal/models/message"
	"context"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/sirupsen/logrus"
)

const Port = 10578

type chatServer struct {
	listeners []message.MessageHandler_SubscribeMessageServer
}

func NewServer() *chatServer {
	return &chatServer{
		listeners: []message.MessageHandler_SubscribeMessageServer{},
	}
}

func (c *chatServer) PublishMessage(ctx context.Context, request *message.MessageRequest) (*flatbuffers.Builder, error) {
	logrus.Debugf("publishing message: [%s->%s]: %s", string(request.Source()), string(request.Destination()), string(request.Body()))

	for _, s := range c.listeners {
		err := s.Send(buildMessageFromRequest(request))

		if err != nil {
			c.removeListener(s)
		}
	}
	return buildMessageFromRequest(request), nil
}

func (c *chatServer) SubscribeMessage(request *message.MessageRequest, stream message.MessageHandler_SubscribeMessageServer) error {
	logrus.Debugf("rcvd a sub to: [%s] from [%s]", string(request.Destination()), string(request.Source()))

	c.addListener(stream)
	<-stream.Context().Done()

	return stream.Context().Err()
}

func (c *chatServer) addListener(stream message.MessageHandler_SubscribeMessageServer) {
	c.listeners = append(c.listeners, stream)
}

func (c *chatServer) removeListener(stream message.MessageHandler_SubscribeMessageServer) {
	c.listeners = append(c.listeners, stream)
}
