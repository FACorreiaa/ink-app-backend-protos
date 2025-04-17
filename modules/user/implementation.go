package studio

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/user/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn

	// Separate clients for each service
	client generated.UserServiceClient
}

// Ensure Broker implements both service interfaces
var (
	_ generated.UserServiceClient = (*Broker)(nil)
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
	// Initialize both clients with the same connection
	b.client = generated.NewUserServiceClient(b.conn)
	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// --- StudioService implementation ---

// GetAllUsers implements generated.Studioclient
func (b *Broker) GetUsers(ctx context.Context, in *generated.GetUsersReq, opts ...grpc.CallOption) (*generated.GetUsersRes, error) {
	return b.client.GetUsers(ctx, in, opts...)
}

// GetUserByID implements generated.Studioclient
func (b *Broker) GetUserByID(ctx context.Context, in *generated.GetUserByIDReq, opts ...grpc.CallOption) (*generated.GetUserByIDRes, error) {
	return b.client.GetUserByID(ctx, in, opts...)
}

// DeleteUser implements generated.Studioclient
func (b *Broker) DeleteUser(ctx context.Context, in *generated.DeleteUserReq, opts ...grpc.CallOption) (*generated.DeleteUserRes, error) {
	return b.client.DeleteUser(ctx, in, opts...)
}

// UpdateUser implements generated.Studioclient
func (b *Broker) UpdateUser(ctx context.Context, in *generated.UpdateUserReq, opts ...grpc.CallOption) (*generated.UpdateUserRes, error) {
	return b.client.UpdateUser(ctx, in, opts...)
}

// InsertUser implements generated.Studioclient
func (b *Broker) InsertUser(ctx context.Context, in *generated.InsertUserReq, opts ...grpc.CallOption) (*generated.InsertUserRes, error) {
	return b.client.InsertUser(ctx, in, opts...)
}

// GetUserByEmail implements generated.Studioclient
func (b *Broker) GetUserByEmail(ctx context.Context, in *generated.GetUserByEmailReq, opts ...grpc.CallOption) (*generated.GetUserByEmailRes, error) {
	return b.client.GetUserByEmail(ctx, in, opts...)
}

// GetUserByUsername implements generated.Studioclient
func (b *Broker) GetUserByUsername(ctx context.Context, in *generated.GetUserByUsernameReq, opts ...grpc.CallOption) (*generated.GetUserByUsernameRes, error) {
	return b.client.GetUserByUsername(ctx, in, opts...)
}
