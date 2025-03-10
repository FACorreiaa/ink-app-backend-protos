package appointment

import (
	"context"
	"errors"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/appointment/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
	"google.golang.org/grpc"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.AppointmentServiceClient
}

var (
	_ generated.AppointmentServiceClient = (*Broker)(nil)
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
	b.client = generated.NewAppointmentServiceClient(b.conn)
	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// ApproveAppointmentRequest implements appointment.AppointmentServiceClient.
func (b *Broker) ApproveAppointmentRequest(ctx context.Context, in *generated.ApproveAppointmentRequestRequest, opts ...grpc.CallOption) (*generated.ApproveAppointmentRequestResponse, error) {
	return b.ApproveAppointmentRequest(ctx, in, opts...)
}

// CancelAppointment implements appointment.AppointmentServiceClient.
func (b *Broker) CancelAppointment(ctx context.Context, in *generated.CancelAppointmentRequest, opts ...grpc.CallOption) (*generated.CancelAppointmentResponse, error) {
	return b.CancelAppointment(ctx, in, opts...)
}

// CreateAppointment implements appointment.AppointmentServiceClient.
func (b *Broker) CreateAppointment(ctx context.Context, in *generated.CreateAppointmentRequest, opts ...grpc.CallOption) (*generated.CreateAppointmentResponse, error) {
	return b.CreateAppointment(ctx, in, opts...)
}

// GetAppointment implements appointment.AppointmentServiceClient.
func (b *Broker) GetAppointment(ctx context.Context, in *generated.GetAppointmentRequest, opts ...grpc.CallOption) (*generated.GetAppointmentResponse, error) {
	return b.GetAppointment(ctx, in, opts...)
}

// GetAvailability implements appointment.AppointmentServiceClient.
func (b *Broker) GetAvailability(ctx context.Context, in *generated.GetAvailabilityRequest, opts ...grpc.CallOption) (*generated.GetAvailabilityResponse, error) {
	return b.GetAvailability(ctx, in, opts...)
}

// ListAppointments implements appointment.AppointmentServiceClient.
func (b *Broker) ListAppointments(ctx context.Context, in *generated.ListAppointmentsRequest, opts ...grpc.CallOption) (*generated.ListAppointmentsResponse, error) {
	return b.ListAppointments(ctx, in, opts...)
}

// ListAvailableTimeSlots implements appointment.AppointmentServiceClient.
func (b *Broker) ListAvailableTimeSlots(ctx context.Context, in *generated.ListAvailableTimeSlotsRequest, opts ...grpc.CallOption) (*generated.ListAvailableTimeSlotsResponse, error) {
	return b.ListAvailableTimeSlots(ctx, in, opts...)
}

// RejectAppointmentRequest implements appointment.AppointmentServiceClient.
func (b *Broker) RejectAppointmentRequest(ctx context.Context, in *generated.RejectAppointmentRequestRequest, opts ...grpc.CallOption) (*generated.RejectAppointmentRequestResponse, error) {
	return b.client.RejectAppointmentRequest(ctx, in, opts...)
}

// RequestAppointment implements appointment.AppointmentServiceClient.
func (b *Broker) RequestAppointment(ctx context.Context, in *generated.RequestAppointmentRequest, opts ...grpc.CallOption) (*generated.RequestAppointmentResponse, error) {
	return b.RequestAppointment(ctx, in, opts...)
}

// RescheduleAppointment implements appointment.AppointmentServiceClient.
func (b *Broker) RescheduleAppointment(ctx context.Context, in *generated.RescheduleAppointmentRequest, opts ...grpc.CallOption) (*generated.RescheduleAppointmentResponse, error) {
	return b.RescheduleAppointment(ctx, in, opts...)
}

// SendAppointmentReminder implements appointment.AppointmentServiceClient.
func (b *Broker) SendAppointmentReminder(ctx context.Context, in *generated.SendAppointmentReminderRequest, opts ...grpc.CallOption) (*generated.SendAppointmentReminderResponse, error) {
	return b.SendAppointmentReminder(ctx, in, opts...)
}

// SetAvailability implements appointment.AppointmentServiceClient.
func (b *Broker) SetAvailability(ctx context.Context, in *generated.SetAvailabilityRequest, opts ...grpc.CallOption) (*generated.SetAvailabilityResponse, error) {
	return b.SetAvailability(ctx, in, opts...)
}

// UpdateAppointment implements appointment.AppointmentServiceClient.
func (b *Broker) UpdateAppointment(ctx context.Context, in *generated.UpdateAppointmentRequest, opts ...grpc.CallOption) (*generated.UpdateAppointmentResponse, error) {
	return b.UpdateAppointment(ctx, in, opts...)
}
