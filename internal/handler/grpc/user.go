package grpc

import (
	"github.com/s3rzh/go-grpc-user-service/internal/service"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"golang.org/x/net/context"
)

type UserManagementServer struct {
	Service *service.Service
	api.UnimplementedUserManagementServer
}

func NewUserManagementServer(service *service.Service) *UserManagementServer {
	return &UserManagementServer{
		Service: service,
	}
}

func (s *UserManagementServer) CreateUser(ctx context.Context, r *api.User) (*api.UserResponse, error) {
	err := r.Validate()
	if err != nil {
		return &api.UserResponse{Msg: err.Error()}, nil
	}

	s.Service.CreateUser(ctx, r)

	return &api.UserResponse{Msg: "user 18 added!"}, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, r *api.EmptyParams) (*api.UsersResponse, error) {
	return &api.UsersResponse{}, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, r *api.UserEmail) (*api.UserResponse, error) {
	err := r.Validate()
	if err != nil {
		return &api.UserResponse{Msg: err.Error()}, nil
	}
	return &api.UserResponse{}, nil
}
