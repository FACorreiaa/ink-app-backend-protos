package portfolio

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	generated "github.com/FACorreiaa/ink-app-backend-protos/modules/portfolio/generated"
	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.PortfolioServiceClient
}

// AssignDesignToCategory implements portfolio.PortfolioServiceClient.
func (b *Broker) AssignDesignToCategory(ctx context.Context, in *generated.AssignDesignToCategoryRequest, opts ...grpc.CallOption) (*generated.AssignDesignToCategoryResponse, error) {
	return b.client.AssignDesignToCategory(ctx, in, opts...)
}

// CreateDesignCategory implements portfolio.PortfolioServiceClient.
func (b *Broker) CreateDesignCategory(ctx context.Context, in *generated.CreateDesignCategoryRequest, opts ...grpc.CallOption) (*generated.CreateDesignCategoryResponse, error) {
	return b.client.CreateDesignCategory(ctx, in, opts...)
}

// DeleteDesign implements portfolio.PortfolioServiceClient.
func (b *Broker) DeleteDesign(ctx context.Context, in *generated.DeleteDesignRequest, opts ...grpc.CallOption) (*generated.DeleteDesignResponse, error) {
	return b.client.DeleteDesign(ctx, in, opts...)
}

// GetDesign implements portfolio.PortfolioServiceClient.
func (b *Broker) GetDesign(ctx context.Context, in *generated.GetDesignRequest, opts ...grpc.CallOption) (*generated.GetDesignResponse, error) {
	return b.client.GetDesign(ctx, in, opts...)
}

// GetSharedDesignsForCustomer implements portfolio.PortfolioServiceClient.
func (b *Broker) GetSharedDesignsForCustomer(ctx context.Context, in *generated.GetSharedDesignsForCustomerRequest, opts ...grpc.CallOption) (*generated.GetSharedDesignsForCustomerResponse, error) {
	return b.client.GetSharedDesignsForCustomer(ctx, in, opts...)
}

// ListDesignCategories implements portfolio.PortfolioServiceClient.
func (b *Broker) ListDesignCategories(ctx context.Context, in *generated.ListDesignCategoriesRequest, opts ...grpc.CallOption) (*generated.ListDesignCategoriesResponse, error) {
	return b.client.ListDesignCategories(ctx, in, opts...)
}

// ListDesigns implements portfolio.PortfolioServiceClient.
func (b *Broker) ListDesigns(ctx context.Context, in *generated.ListDesignsRequest, opts ...grpc.CallOption) (*generated.ListDesignsResponse, error) {
	return b.client.ListDesigns(ctx, in, opts...)
}

// ShareDesignWithCustomer implements portfolio.PortfolioServiceClient.
func (b *Broker) ShareDesignWithCustomer(ctx context.Context, in *generated.ShareDesignWithCustomerRequest, opts ...grpc.CallOption) (*generated.ShareDesignWithCustomerResponse, error) {
	return b.client.ShareDesignWithCustomer(ctx, in, opts...)
}

// UpdateDesign implements portfolio.PortfolioServiceClient.
func (b *Broker) UpdateDesign(ctx context.Context, in *generated.UpdateDesignRequest, opts ...grpc.CallOption) (*generated.UpdateDesignResponse, error) {
	return b.client.UpdateDesign(ctx, in, opts...)
}

// UploadDesign implements portfolio.PortfolioServiceClient.
func (b *Broker) UploadDesign(ctx context.Context, in *generated.UploadDesignRequest, opts ...grpc.CallOption) (*generated.UploadDesignResponse, error) {
	return b.client.UploadDesign(ctx, in, opts...)
}

var (
	_ generated.PortfolioServiceClient = (*Broker)(nil)
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
	b.client = generated.NewPortfolioServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}
