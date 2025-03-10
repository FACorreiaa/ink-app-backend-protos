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
	client     generated.StudioServiceClient
}

// AddStaffMember implements studio.StudioServiceClient.
func (b *Broker) AddStaffMember(ctx context.Context, in *generated.AddStaffMemberRequest, opts ...grpc.CallOption) (*generated.AddStaffMemberResponse, error) {
	return b.client.AddStaffMember(ctx, in, opts...)
}

// CreateStudio implements studio.StudioServiceClient.
func (b *Broker) CreateStudio(ctx context.Context, in *generated.CreateStudioRequest, opts ...grpc.CallOption) (*generated.CreateStudioResponse, error) {
	return b.client.CreateStudio(ctx, in, opts...)
}

// GetStaffPermissions implements studio.StudioServiceClient.
func (b *Broker) GetStaffPermissions(ctx context.Context, in *generated.GetStaffPermissionsRequest, opts ...grpc.CallOption) (*generated.GetStaffPermissionsResponse, error) {
	return b.client.GetStaffPermissions(ctx, in, opts...)
}

// GetStudio implements studio.StudioServiceClient.
func (b *Broker) GetStudio(ctx context.Context, in *generated.GetStudioRequest, opts ...grpc.CallOption) (*generated.GetStudioResponse, error) {
	return b.client.GetStudio(ctx, in, opts...)
}

// ListStaffMembers implements studio.StudioServiceClient.
func (b *Broker) ListStaffMembers(ctx context.Context, in *generated.ListStaffMembersRequest, opts ...grpc.CallOption) (*generated.ListStaffMembersResponse, error) {
	return b.client.ListStaffMembers(ctx, in, opts...)
}

// RemoveStaffMember implements studio.StudioServiceClient.
func (b *Broker) RemoveStaffMember(ctx context.Context, in *generated.RemoveStaffMemberRequest, opts ...grpc.CallOption) (*generated.RemoveStaffMemberResponse, error) {
	return b.client.RemoveStaffMember(ctx, in, opts...)
}

// SetStaffPermissions implements studio.StudioServiceClient.
func (b *Broker) SetStaffPermissions(ctx context.Context, in *generated.SetStaffPermissionsRequest, opts ...grpc.CallOption) (*generated.SetStaffPermissionsResponse, error) {
	return b.client.SetStaffPermissions(ctx, in, opts...)
}

// UpdateStaffMember implements studio.StudioServiceClient.
func (b *Broker) UpdateStaffMember(ctx context.Context, in *generated.UpdateStaffMemberRequest, opts ...grpc.CallOption) (*generated.UpdateStaffMemberResponse, error) {
	return b.client.UpdateStaffMember(ctx, in, opts...)
}

// UpdateStudio implements studio.StudioServiceClient.
func (b *Broker) UpdateStudio(ctx context.Context, in *generated.UpdateStudioRequest, opts ...grpc.CallOption) (*generated.UpdateStudioResponse, error) {
	return b.client.UpdateStudio(ctx, in, opts...)
}

var (
	_ generated.StudioServiceClient = (*Broker)(nil)
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
	b.client = generated.NewStudioServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}
