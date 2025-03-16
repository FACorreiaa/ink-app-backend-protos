package studio

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/studio/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn

	// Separate clients for each service
	studioClient generated.StudioServiceClient
	authClient   generated.StudioAuthClient
}

// Ensure Broker implements both service interfaces
var (
	_ generated.StudioServiceClient = (*Broker)(nil)
	_ generated.StudioAuthClient    = (*Broker)(nil)
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
	b.studioClient = generated.NewStudioServiceClient(b.conn)
	b.authClient = generated.NewStudioAuthClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// --- StudioService implementation ---

// AddStudioUser implements generated.StudioServiceClient
func (b *Broker) AddStudioUser(ctx context.Context, in *generated.AddStudioUserRequest, opts ...grpc.CallOption) (*generated.AddStudioUserResponse, error) {
	return b.studioClient.AddStudioUser(ctx, in, opts...)
}

// ListStudios implements generated.StudioServiceClient
func (b *Broker) ListStudios(ctx context.Context, in *generated.ListStudiosRequest, opts ...grpc.CallOption) (*generated.ListStudiosResponse, error) {
	return b.studioClient.ListStudios(ctx, in, opts...)
}

// RemoveStudioUser implements generated.StudioServiceClient
func (b *Broker) RemoveStudioUser(ctx context.Context, in *generated.RemoveStudioUserRequest, opts ...grpc.CallOption) (*generated.RemoveStudioUserResponse, error) {
	return b.studioClient.RemoveStudioUser(ctx, in, opts...)
}

// UpdateStudioUser implements generated.StudioServiceClient
func (b *Broker) UpdateStudioUser(ctx context.Context, in *generated.UpdateStudioUserRequest, opts ...grpc.CallOption) (*generated.UpdateStudioUserResponse, error) {
	return b.studioClient.UpdateStudioUser(ctx, in, opts...)
}

// AddStaffMember implements generated.StudioServiceClient
func (b *Broker) AddStaffMember(ctx context.Context, in *generated.AddStaffMemberRequest, opts ...grpc.CallOption) (*generated.AddStaffMemberResponse, error) {
	return b.studioClient.AddStaffMember(ctx, in, opts...)
}

// CreateStudio implements generated.StudioServiceClient
func (b *Broker) CreateStudio(ctx context.Context, in *generated.CreateStudioRequest, opts ...grpc.CallOption) (*generated.CreateStudioResponse, error) {
	return b.studioClient.CreateStudio(ctx, in, opts...)
}

// GetStaffPermissions implements generated.StudioServiceClient
func (b *Broker) GetStaffPermissions(ctx context.Context, in *generated.GetStaffPermissionsRequest, opts ...grpc.CallOption) (*generated.GetStaffPermissionsResponse, error) {
	return b.studioClient.GetStaffPermissions(ctx, in, opts...)
}

// ListStaffMembers implements generated.StudioServiceClient
func (b *Broker) ListStaffMembers(ctx context.Context, in *generated.ListStaffMembersRequest, opts ...grpc.CallOption) (*generated.ListStaffMembersResponse, error) {
	return b.studioClient.ListStaffMembers(ctx, in, opts...)
}

// RemoveStaffMember implements generated.StudioServiceClient
func (b *Broker) RemoveStaffMember(ctx context.Context, in *generated.RemoveStaffMemberRequest, opts ...grpc.CallOption) (*generated.RemoveStaffMemberResponse, error) {
	return b.studioClient.RemoveStaffMember(ctx, in, opts...)
}

// SetStaffPermissions implements generated.StudioServiceClient
func (b *Broker) SetStaffPermissions(ctx context.Context, in *generated.SetStaffPermissionsRequest, opts ...grpc.CallOption) (*generated.SetStaffPermissionsResponse, error) {
	return b.studioClient.SetStaffPermissions(ctx, in, opts...)
}

// UpdateStaffMember implements generated.StudioServiceClient
func (b *Broker) UpdateStaffMember(ctx context.Context, in *generated.UpdateStaffMemberRequest, opts ...grpc.CallOption) (*generated.UpdateStaffMemberResponse, error) {
	return b.studioClient.UpdateStaffMember(ctx, in, opts...)
}

// UpdateStudio implements generated.StudioServiceClient
func (b *Broker) UpdateStudio(ctx context.Context, in *generated.UpdateStudioRequest, opts ...grpc.CallOption) (*generated.UpdateStudioResponse, error) {
	return b.studioClient.UpdateStudio(ctx, in, opts...)
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
func (b *Broker) Logout(ctx context.Context, in *generated.LogoutRequest, opts ...grpc.CallOption) (*generated.LogoutResponse, error) {
	return b.authClient.Logout(ctx, in, opts...)
}

func (b *Broker) ValidateSession(ctx context.Context, in *generated.ValidateSessionRequest, opts ...grpc.CallOption) (*generated.ValidateSessionResponse, error) {
	return b.authClient.ValidateSession(ctx, in, opts...)
}

// ChangePassword implements generated.StudioAuthClient
func (b *Broker) ChangePassword(ctx context.Context, in *generated.ChangePasswordRequest, opts ...grpc.CallOption) (*generated.ChangePasswordResponse, error) {
	return b.authClient.ChangePassword(ctx, in, opts...)
}

// ChangeEmail implements generated.StudioAuthClient
func (b *Broker) ChangeEmail(ctx context.Context, in *generated.ChangeEmailRequest, opts ...grpc.CallOption) (*generated.ChangeEmailResponse, error) {
	return b.authClient.ChangeEmail(ctx, in, opts...)
}

// GetAllUsers implements generated.StudioAuthClient
func (b *Broker) GetAllUsers(ctx context.Context, in *generated.GetAllUsersRequest, opts ...grpc.CallOption) (*generated.GetAllUsersResponse, error) {
	return b.authClient.GetAllUsers(ctx, in, opts...)
}

// GetUserByID implements generated.StudioAuthClient
func (b *Broker) GetUserByID(ctx context.Context, in *generated.GetUserByIDRequest, opts ...grpc.CallOption) (*generated.GetUserByIDResponse, error) {
	return b.authClient.GetUserByID(ctx, in, opts...)
}

// DeleteUser implements generated.StudioAuthClient
func (b *Broker) DeleteUser(ctx context.Context, in *generated.DeleteUserRequest, opts ...grpc.CallOption) (*generated.DeleteUserResponse, error) {
	return b.authClient.DeleteUser(ctx, in, opts...)
}

// UpdateUser implements generated.StudioAuthClient
func (b *Broker) UpdateUser(ctx context.Context, in *generated.UpdateUserRequest, opts ...grpc.CallOption) (*generated.UpdateUserResponse, error) {
	return b.authClient.UpdateUser(ctx, in, opts...)
}

// InsertUser implements generated.StudioAuthClient
func (b *Broker) InsertUser(ctx context.Context, in *generated.InsertUserRequest, opts ...grpc.CallOption) (*generated.InsertUserResponse, error) {
	return b.authClient.InsertUser(ctx, in, opts...)
}

// RefreshToken implements generated.StudioAuthClient
func (b *Broker) RefreshToken(ctx context.Context, in *generated.RefreshTokenRequest, opts ...grpc.CallOption) (*generated.TokenResponse, error) {
	return b.authClient.RefreshToken(ctx, in, opts...)
}
