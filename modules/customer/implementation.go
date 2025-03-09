package client

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	Customer "github.com/FACorreiaa/ink-app-backend-protos/modules/customer/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     Customer.CustomerServiceClient
}

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
	b.client = Customer.NewCustomerServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// Implement the CustomerServiceClient interface methods
func (b *Broker) CreateCustomer(ctx context.Context, in *Customer.CreateCustomerRequest, opts ...grpc.CallOption) (*Customer.CreateCustomerResponse, error) {
	return b.client.CreateCustomer(ctx, in, opts...)
}

func (b *Broker) GetCustomer(ctx context.Context, in *Customer.GetCustomerRequest, opts ...grpc.CallOption) (*Customer.GetCustomerResponse, error) {
	return b.client.GetCustomer(ctx, in, opts...)
}

func (b *Broker) UpdateCustomer(ctx context.Context, in *Customer.UpdateCustomerRequest, opts ...grpc.CallOption) (*Customer.UpdateCustomerResponse, error) {
	return b.client.UpdateCustomer(ctx, in, opts...)
}

func (b *Broker) DeleteCustomer(ctx context.Context, in *Customer.DeleteCustomerRequest, opts ...grpc.CallOption) (*Customer.DeleteCustomerResponse, error) {
	return b.client.DeleteCustomer(ctx, in, opts...)
}

func (b *Broker) ListCustomers(ctx context.Context, in *Customer.ListCustomersRequest, opts ...grpc.CallOption) (*Customer.ListCustomersResponse, error) {
	return b.client.ListCustomers(ctx, in, opts...)
}

func (b *Broker) ArchiveCustomer(ctx context.Context, in *Customer.ArchiveCustomerRequest, opts ...grpc.CallOption) (*Customer.ArchiveCustomerResponse, error) {
	return b.client.ArchiveCustomer(ctx, in, opts...)
}

func (b *Broker) GetCustomerHistory(ctx context.Context, in *Customer.GetCustomerHistoryRequest, opts ...grpc.CallOption) (*Customer.GetCustomerHistoryResponse, error) {
	return b.client.GetCustomerHistory(ctx, in, opts...)
}

func (b *Broker) AddCustomerNote(ctx context.Context, in *Customer.AddCustomerNoteRequest, opts ...grpc.CallOption) (*Customer.AddCustomerNoteResponse, error) {
	return b.client.AddCustomerNote(ctx, in, opts...)
}

func (b *Broker) SearchCustomers(ctx context.Context, in *Customer.SearchCustomersRequest, opts ...grpc.CallOption) (*Customer.SearchCustomersResponse, error) {
	return b.client.SearchCustomers(ctx, in, opts...)
}
