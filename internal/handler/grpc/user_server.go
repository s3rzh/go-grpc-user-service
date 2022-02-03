package grpc

import (
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"golang.org/x/net/context"
)

type UserManagementServer struct {
	api.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateUser(ctx context.Context, r *api.User) (*api.UserResponse, error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}

	return &api.UserResponse{}, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, r *api.EmptyParams) (*api.UsersResponse, error) {
	return &api.UsersResponse{}, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, r *api.UserEmail) (*api.UserResponse, error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}
	return &api.UserResponse{}, nil
}
