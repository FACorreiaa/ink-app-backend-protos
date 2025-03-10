package conversation

import (
	"context"
	"errors"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/conversation/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
	"google.golang.org/grpc"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.ConversationServiceClient
}

var (
	_ generated.ConversationServiceClient = (*Broker)(nil)
)

func NewBroker(serverAddr string) (*Broker, error) {
	b := new(Broker)
	b.serverAddr = serverAddr

	if b.serverAddr == "" {
		return nil, errors.New("null routed upstread host")
	}

	return b, nil
}

func (b *Broker) NewConnection() (*grpc.ClientConn, error) {
	conn, err := utils.NewConnection(b.serverAddr)
	if err != nil {
		return nil, errors.New("could not open connection")
	}

	b.conn = conn
	b.client = generated.NewConversationServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// ArchiveConversation implements conversation.ConversationServiceClient.
func (b *Broker) ArchiveConversation(ctx context.Context, in *generated.ArchiveConversationRequest, opts ...grpc.CallOption) (*generated.ArchiveConversationResponse, error) {
	return b.ArchiveConversation(ctx, in, opts...)
}

// CreateConversation implements conversation.ConversationServiceClient.
func (b *Broker) CreateConversation(ctx context.Context, in *generated.CreateConversationRequest, opts ...grpc.CallOption) (*generated.CreateConversationResponse, error) {
	return b.CreateConversation(ctx, in, opts...)
}

// GetConversation implements conversation.ConversationServiceClient.
func (b *Broker) GetConversation(ctx context.Context, in *generated.GetConversationRequest, opts ...grpc.CallOption) (*generated.GetConversationResponse, error) {
	return b.GetConversation(ctx, in, opts...)
}

// ListConversations implements conversation.ConversationServiceClient.
func (b *Broker) ListConversations(ctx context.Context, in *generated.ListConversationsRequest, opts ...grpc.CallOption) (*generated.ListConversationsResponse, error) {
	return b.ListConversations(ctx, in, opts...)
}

// MarkAsRead implements conversation.ConversationServiceClient.
func (b *Broker) MarkAsRead(ctx context.Context, in *generated.MarkAsReadRequest, opts ...grpc.CallOption) (*generated.MarkAsReadResponse, error) {
	return b.MarkAsRead(ctx, in, opts...)
}

// SendMessage implements conversation.ConversationServiceClient.
func (b *Broker) SendMessage(ctx context.Context, in *generated.SendMessageRequest, opts ...grpc.CallOption) (*generated.SendMessageResponse, error) {
	return b.SendMessage(ctx, in, opts...)
}
