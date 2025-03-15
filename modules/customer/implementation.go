package customer

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/customer/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.CustomerServiceClient
	authClient generated.CustomerAuthClient
}

func NewBroker(serverAddr string) (*Broker, error) {
	b := new(Broker)
	b.serverAddr = serverAddr

	if b.serverAddr == "" {
		return nil, errors.New("null routed upstream host")
	}

	return b, nil
}

// Ensure Broker implements both service interfaces
var (
	_ generated.CustomerServiceClient = (*Broker)(nil)
	_ generated.CustomerAuthClient    = (*Broker)(nil)
)

func (b *Broker) NewConnection() (*grpc.ClientConn, error) {
	conn, err := utils.NewConnection(b.serverAddr)
	if err != nil {
		return nil, errors.New("could not open connection")
	}

	b.conn = conn
	b.client = generated.NewCustomerServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// Implement the CustomerServiceClient interface methods
func (b *Broker) CreateCustomer(ctx context.Context, in *generated.CreateCustomerRequest, opts ...grpc.CallOption) (*generated.CreateCustomerResponse, error) {
	return b.client.CreateCustomer(ctx, in, opts...)
}

func (b *Broker) GetCustomer(ctx context.Context, in *generated.GetCustomerRequest, opts ...grpc.CallOption) (*generated.GetCustomerResponse, error) {
	return b.client.GetCustomer(ctx, in, opts...)
}

func (b *Broker) UpdateCustomer(ctx context.Context, in *generated.UpdateCustomerRequest, opts ...grpc.CallOption) (*generated.UpdateCustomerResponse, error) {
	return b.client.UpdateCustomer(ctx, in, opts...)
}

func (b *Broker) DeleteCustomer(ctx context.Context, in *generated.DeleteCustomerRequest, opts ...grpc.CallOption) (*generated.DeleteCustomerResponse, error) {
	return b.client.DeleteCustomer(ctx, in, opts...)
}

func (b *Broker) ListCustomers(ctx context.Context, in *generated.ListCustomersRequest, opts ...grpc.CallOption) (*generated.ListCustomersResponse, error) {
	return b.client.ListCustomers(ctx, in, opts...)
}

func (b *Broker) ArchiveCustomer(ctx context.Context, in *generated.ArchiveCustomerRequest, opts ...grpc.CallOption) (*generated.ArchiveCustomerResponse, error) {
	return b.client.ArchiveCustomer(ctx, in, opts...)
}

func (b *Broker) GetCustomerHistory(ctx context.Context, in *generated.GetCustomerHistoryRequest, opts ...grpc.CallOption) (*generated.GetCustomerHistoryResponse, error) {
	return b.client.GetCustomerHistory(ctx, in, opts...)
}

func (b *Broker) AddCustomerNote(ctx context.Context, in *generated.AddCustomerNoteRequest, opts ...grpc.CallOption) (*generated.AddCustomerNoteResponse, error) {
	return b.client.AddCustomerNote(ctx, in, opts...)
}

func (b *Broker) SearchCustomers(ctx context.Context, in *generated.SearchCustomersRequest, opts ...grpc.CallOption) (*generated.SearchCustomersResponse, error) {
	return b.client.SearchCustomers(ctx, in, opts...)
}

// --- StudioAuth implementation ---

// Register implements generated.StudioAuthClient
func (b *Broker) Register(ctx context.Context, in *generated.RegisterRequest, opts ...grpc.CallOption) (*generated.RegisterResponse, error) {
	return b.authClient.Register(ctx, in, opts...)
}

// Login implements generated.StudioAuthClient
func (b *Broker) Login(ctx context.Context, in *generated.LoginRequest, opts ...grpc.CallOption) (*generated.LoginResponse, error) {
	return b.authClient.Login(ctx, in, opts...)
}

// Logout implements generated.StudioAuthClient
func (b *Broker) Logout(ctx context.Context, in *generated.NilReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.authClient.Logout(ctx, in, opts...)
}

// ChangePassword implements generated.StudioAuthClient
func (b *Broker) ChangePassword(ctx context.Context, in *generated.ChangePasswordRequest, opts ...grpc.CallOption) (*generated.ChangePasswordResponse, error) {
	return b.authClient.ChangePassword(ctx, in, opts...)
}

// ChangeEmail implements generated.StudioAuthClient
func (b *Broker) ChangeEmail(ctx context.Context, in *generated.ChangeEmailRequest, opts ...grpc.CallOption) (*generated.ChangeEmailResponse, error) {
	return b.authClient.ChangeEmail(ctx, in, opts...)
}

// UpdateUser implements generated.StudioAuthClient
func (b *Broker) UpdateUser(ctx context.Context, in *generated.UpdateUserRequest, opts ...grpc.CallOption) (*generated.UpdateUserResponse, error) {
	return b.authClient.UpdateUser(ctx, in, opts...)
}

// RefreshToken implements generated.StudioAuthClient
func (b *Broker) RefreshToken(ctx context.Context, in *generated.RefreshTokenRequest, opts ...grpc.CallOption) (*generated.TokenResponse, error) {
	return b.authClient.RefreshToken(ctx, in, opts...)
}
