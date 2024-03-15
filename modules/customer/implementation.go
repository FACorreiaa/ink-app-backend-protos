package customer

import (
	"context"
	"errors"

	"github.com/highly-regarded/grpc/core"
	"github.com/highly-regarded/grpc/modules/customer/generated"
	"github.com/highly-regarded/grpc/utils"
	"google.golang.org/grpc"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.CustomerClient
}

var (
	_ generated.CustomerClient = (*Broker)(nil)
	_ core.Broker              = (*Broker)(nil)
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
	b.client = generated.NewCustomerClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetCustomer(ctx context.Context, in *generated.GetCustomerReq, opts ...grpc.CallOption) (*generated.GetCustomerRes, error) {
	return b.client.GetCustomer(ctx, in, opts...)
}

func (b *Broker) CreateCustomer(ctx context.Context, in *generated.CreateCustomerReq, opts ...grpc.CallOption) (*generated.CreateCustomerRes, error) {
	return b.client.CreateCustomer(ctx, in, opts...)
}

func (b *Broker) UpdateCustomer(ctx context.Context, in *generated.UpdateCustomerReq, opts ...grpc.CallOption) (*generated.UpdateCustomerRes, error) {
	return b.client.UpdateCustomer(ctx, in, opts...)
}

func (b *Broker) DeleteCustomer(ctx context.Context, in *generated.DeleteCustomerReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteCustomer(ctx, in, opts...)
}
