package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
	Scope  string `json:"scope"`
	jwt.RegisteredClaims
}

var JwtSecretKey = []byte("your-secret-key")
var JwtRefreshSecretKey = []byte("your-refresh-key")

// CustomerService implements the Customer gRPC server
//type CustomerService struct {
//	pb.UnimplementedCustomerServer
//	ctx    context.Context
//	pgpool *pgxpool.Pool
//	redis  *redis.Client
//}
//
//// NewCustomerService creates a new CustomerService
//func NewCustomerService(ctx context.Context, db *pgxpool.Pool, redis *redis.Client) *CustomerService {
//	return &CustomerService{ctx: ctx, pgpool: db, redis: redis}
//}
//
//func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerReq) (*pb.GetCustomerRes, error) {
//	// Implementation of GetCustomer
//	return &pb.GetCustomerRes{}, nil
//}
//
//func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerReq) (*pb.CreateCustomerRes, error) {
//	return &pb.CreateCustomerRes{}, nil
//}
//
//func (s *CustomerService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerReq) (*pb.UpdateCustomerRes, error) {
//	// Implementation of UpdateCustomer
//	return &pb.UpdateCustomerRes{}, nil
//}
//
//func (s *CustomerService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerReq) (*pb.NilRes, error) {
//	// Implementation of DeleteCustomer
//	return &pb.NilRes{}, nil
//}

// AuthService

// Calculator service

// Meal service
