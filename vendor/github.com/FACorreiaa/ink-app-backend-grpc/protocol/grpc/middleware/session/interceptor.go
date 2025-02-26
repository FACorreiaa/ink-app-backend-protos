package session

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/FACorreiaa/ink-app-backend-grpc/internal/domain"
)

// Define your secret key for signing tokens

// Claims struct

func InterceptorSession() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		unauthenticatedMethods := map[string]bool{
			"/inkMe.auth.Auth/Register":    true,
			"/inkMe.auth.Auth/Login":       true,
			"/inkMe.auth.Auth/GetAllUsers": true,
		}
		if unauthenticatedMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing context metadata")
		}

		authHeader := md["authorization"]
		//if len(authHeader) == 0 || len(authHeader[0]) < 8 || authHeader[0][:7] != "Bearer " {
		//	return nil, status.Error(codes.Unauthenticated, "missing or invalid auth token")
		//}
		//
		//tokenString := authHeader[0][7:]
		println(authHeader)
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing or invalid auth token")
		}

		tokenString := authHeader[0]

		claims := &domain.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return domain.JwtSecretKey, nil
		})
		if err != nil || !token.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}

		ctx = context.WithValue(ctx, "userID", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)

		return handler(ctx, req)
	}
}

func hasPermission(userPermissions []string, requiredPermission string) bool {
	for _, perm := range userPermissions {
		if perm == requiredPermission {
			return true
		}
	}
	return false
}
