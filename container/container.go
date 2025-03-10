package container

import (
	auth "github.com/FACorreiaa/ink-app-backend-protos/modules/auth/generated"
	customer "github.com/FACorreiaa/ink-app-backend-protos/modules/customer/generated"

	"github.com/FACorreiaa/ink-app-backend-protos/utils"
)

// Brokers is the encapsulation for all grpc clients. We use it this way so
// that we can:
//
// 1. Mock remotes in tests
// 2. Have a single point of change when contracts change
// 3. Use a common logic route for onboarding
//
// Ensure that you're using the interface type here and not the implementation
type Brokers struct {
	Customer       customer.CustomerServiceClient
	Auth           auth.AuthClient
	TransportUtils *utils.TransportUtils
}

// NewBrokers creates a common container instance for use with all cluster sevices
func NewBrokers(transportUtils *utils.TransportUtils) *Brokers {
	if transportUtils == nil {
		return nil
	}

	brokers := new(Brokers)
	brokers.TransportUtils = transportUtils
	utils.Transport = transportUtils

	return brokers
}
