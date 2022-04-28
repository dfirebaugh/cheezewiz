package chatservice

import (
	"cheezewiz/config"
	"cheezewiz/internal/models/message"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type chatClient struct {
	response string
}

func NewClient() *chatClient {
	return &chatClient{}
}

func withCancel(client message.MessageHandlerClient, _ context.Context) (message.MessageHandlerClient, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return client, ctx, cancel
}

func (*chatClient) getMessageHandlerClient(cc *grpc.ClientConn) (message.MessageHandlerClient, context.Context) {
	return message.NewMessageHandlerClient(cc), context.Background()
}

func (*chatClient) getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Get().Server, config.Get().Port), grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{})))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	return conn
}

func (c *chatClient) PublishMessage(source string, destination string, content string) string {
	conn := c.getConnection()
	defer conn.Close()

	client, ctx, cancel := withCancel(c.getMessageHandlerClient(conn))
	builder := buildMessage(source, destination, content)

	defer cancel()

	request, err := client.PublishMessage(ctx, builder, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Fatalf("%v.SendMessage(_) = _, %v: ", client, err)
	}

	return fmt.Sprintf("[%s->%s]: %s", request.Source(), request.Destination(), request.Body())
}

func (c *chatClient) SubscribeMessage(source string, destination string) error {
	conn := c.getConnection()
	// defer conn.Close()

	client, ctx := c.getMessageHandlerClient(conn)
	builder := buildMessage(source, destination, "")

	stream, err := client.SubscribeMessage(ctx, builder, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		logrus.Error(err)
		return err
	}

	go c.receive(stream)
	return nil
}

func (c *chatClient) receive(stream message.MessageHandler_SubscribeMessageClient) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			logrus.Info("stream ended")
			break
		}

		if err != nil {
			logrus.Error(err)
			break
		}

		c.response = fmt.Sprintf("[%s]:%s", msg.Source(), msg.Body())
	}
}

func (c *chatClient) GetResponse() string {
	return c.response
}
