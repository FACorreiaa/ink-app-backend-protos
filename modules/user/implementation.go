package user

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/ink-app-backend-protos/modules/user/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.AuthClient
}

var (
	_ generated.AuthClient = (*Broker)(nil)
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
	b.client = generated.NewAuthClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// Implement the AuthClient interface methods

func (b *Broker) Register(ctx context.Context, in *generated.RegisterRequest, opts ...grpc.CallOption) (*generated.RegisterResponse, error) {
	return b.client.Register(ctx, in, opts...)
}

func (b *Broker) Login(ctx context.Context, in *generated.LoginRequest, opts ...grpc.CallOption) (*generated.LoginResponse, error) {
	return b.client.Login(ctx, in, opts...)
}

func (b *Broker) Logout(ctx context.Context, in *generated.NilReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.Logout(ctx, in, opts...)
}

func (b *Broker) ChangePassword(ctx context.Context, in *generated.ChangePasswordRequest, opts ...grpc.CallOption) (*generated.ChangePasswordResponse, error) {
	return b.client.ChangePassword(ctx, in, opts...)
}

func (b *Broker) ChangeEmail(ctx context.Context, in *generated.ChangeEmailRequest, opts ...grpc.CallOption) (*generated.ChangeEmailResponse, error) {
	return b.client.ChangeEmail(ctx, in, opts...)
}

func (b *Broker) GetAllUsers(ctx context.Context, in *generated.GetAllUsersRequest, opts ...grpc.CallOption) (*generated.GetAllUsersResponse, error) {
	return b.client.GetAllUsers(ctx, in, opts...)
}

func (b *Broker) GetUserByID(ctx context.Context, in *generated.GetUserByIDRequest, opts ...grpc.CallOption) (*generated.GetUserByIDResponse, error) {
	return b.client.GetUserByID(ctx, in, opts...)
}

func (b *Broker) DeleteUser(ctx context.Context, in *generated.DeleteUserRequest, opts ...grpc.CallOption) (*generated.DeleteUserResponse, error) {
	return b.client.DeleteUser(ctx, in, opts...)
}

func (b *Broker) UpdateUser(ctx context.Context, in *generated.UpdateUserRequest, opts ...grpc.CallOption) (*generated.UpdateUserResponse, error) {
	return b.client.UpdateUser(ctx, in, opts...)
}

func (b *Broker) InsertUser(ctx context.Context, in *generated.InsertUserRequest, opts ...grpc.CallOption) (*generated.InsertUserResponse, error) {
	return b.client.InsertUser(ctx, in, opts...)
}

func (b *Broker) RefreshToken(ctx context.Context, in *generated.RefreshTokenRequest, opts ...grpc.CallOption) (*generated.TokenResponse, error) {
	return b.client.RefreshToken(ctx, in, opts...)
}
