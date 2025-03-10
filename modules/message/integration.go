package message

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/message/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.MessagingServiceClient
}

// CreateMessageTemplate implements message.MessagingServiceClient.
func (b *Broker) CreateMessageTemplate(ctx context.Context, in *generated.CreateMessageTemplateRequest, opts ...grpc.CallOption) (*generated.CreateMessageTemplateResponse, error) {
	return b.CreateMessageTemplate(ctx, in, opts...)
}

// CreateMessageThread implements message.MessagingServiceClient.
func (b *Broker) CreateMessageThread(ctx context.Context, in *generated.CreateMessageThreadRequest, opts ...grpc.CallOption) (*generated.CreateMessageThreadResponse, error) {
	return b.CreateMessageThread(ctx, in, opts...)
}

// GetMessage implements message.MessagingServiceClient.
func (b *Broker) GetMessage(ctx context.Context, in *generated.GetMessageRequest, opts ...grpc.CallOption) (*generated.GetMessageResponse, error) {
	return b.GetMessage(ctx, in, opts...)
}

// GetMessageThread implements message.MessagingServiceClient.
func (b *Broker) GetMessageThread(ctx context.Context, in *generated.GetMessageThreadRequest, opts ...grpc.CallOption) (*generated.GetMessageThreadResponse, error) {
	return b.GetMessageThread(ctx, in, opts...)
}

// ListMessageTemplates implements message.MessagingServiceClient.
func (b *Broker) ListMessageTemplates(ctx context.Context, in *generated.ListMessageTemplatesRequest, opts ...grpc.CallOption) (*generated.ListMessageTemplatesResponse, error) {
	return b.ListMessageTemplates(ctx, in, opts...)
}

// ListMessageThreads implements message.MessagingServiceClient.
func (b *Broker) ListMessageThreads(ctx context.Context, in *generated.ListMessageThreadsRequest, opts ...grpc.CallOption) (*generated.ListMessageThreadsResponse, error) {
	return b.ListMessageThreads(ctx, in, opts...)
}

// ListMessages implements message.MessagingServiceClient.
func (b *Broker) ListMessages(ctx context.Context, in *generated.ListMessagesRequest, opts ...grpc.CallOption) (*generated.ListMessagesResponse, error) {
	return b.ListMessages(ctx, in, opts...)
}

// MarkMessageAsRead implements message.MessagingServiceClient.
func (b *Broker) MarkMessageAsRead(ctx context.Context, in *generated.MarkMessageAsReadRequest, opts ...grpc.CallOption) (*generated.MarkMessageAsReadResponse, error) {
	return b.MarkMessageAsRead(ctx, in, opts...)
}

// SendMessage implements message.MessagingServiceClient.
func (b *Broker) SendMessage(ctx context.Context, in *generated.SendMessageRequest, opts ...grpc.CallOption) (*generated.SendMessageResponse, error) {
	return b.SendMessage(ctx, in, opts...)
}

// SendTemplatedMessage implements message.MessagingServiceClient.
func (b *Broker) SendTemplatedMessage(ctx context.Context, in *generated.SendTemplatedMessageRequest, opts ...grpc.CallOption) (*generated.SendTemplatedMessageResponse, error) {
	return b.SendTemplatedMessage(ctx, in, opts...)
}

// StreamMessages implements message.MessagingServiceClient.
func (b *Broker) StreamMessages(ctx context.Context, in *generated.StreamMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[generated.Message], error) {
	return b.StreamMessages(ctx, in, opts...)
}

var (
	_ generated.MessagingServiceClient = (*Broker)(nil)
)

func NewBroker(serverAddr string) (*Broker, error) {
	b := new(Broker)
	b.serverAddr = serverAddr
	if b.serverAddr == "" {
		return nil, errors.New("null routed upstream host")
	}
	return b, nil
}

func (b *Broker) NewConnection() (*grpc.ClientConn, error) {
	conn, err := utils.NewConnection(b.serverAddr)
	if err != nil {
		return nil, errors.New("could not open connection")
	}
	b.conn = conn
	b.client = generated.NewMessagingServiceClient(b.conn)
	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}
