package grpc

import (
	"fmt"

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

func (s *UserManagementServer) CreateUser(ctx context.Context, u *api.User) (*api.UserResponse, error) {
	var resp api.UserResponse

	err := u.Validate()
	if err != nil {
		resp.Msg = err.Error()
		return &resp, nil
	}

	userId, err := s.Service.CreateUser(ctx, u)
	if err != nil {
		resp.Msg = err.Error()
		return &resp, nil
	}

	resp.Msg = fmt.Sprintf("Added user with id %d", userId)
	return &resp, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, r *api.EmptyParams) (*api.UsersResponse, error) {
	return &api.UsersResponse{}, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, ue *api.UserEmail) (*api.UserResponse, error) {
	var resp api.UserResponse

	err := ue.Validate()
	if err != nil {
		resp.Msg = err.Error()
		return &resp, nil
	}

	err = s.Service.DeleteUser(ctx, ue)
	if err != nil {
		resp.Msg = err.Error()
		return &resp, nil
	}

	resp.Msg = "User with this email was deleted"
	return &resp, nil
}
