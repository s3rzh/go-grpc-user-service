package grpc

import (
	"fmt"

	"github.com/s3rzh/go-grpc-user-service/internal/config"
	"github.com/s3rzh/go-grpc-user-service/internal/service"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"golang.org/x/net/context"
)

type UserManagementServer struct {
	Service  *service.Service
	messages config.Messages
	api.UnimplementedUserManagementServer
}

func NewUserManagementServer(service *service.Service, messages config.Messages) *UserManagementServer {
	return &UserManagementServer{
		Service:  service,
		messages: messages,
	}
}

func (s *UserManagementServer) CreateUser(ctx context.Context, u *api.User) (*api.UserResponse, error) {
	var resp api.UserResponse

	err := u.Validate()
	if err != nil {
		resp.Msg = s.getMessageError(errInputData)
		return &resp, nil
	}

	userId, err := s.Service.CreateUser(ctx, u)
	if err != nil {
		resp.Msg = s.getMessageError(err)
		return &resp, nil
	}

	if userId == 0 {
		resp.Msg = s.getMessageError(errAlreadyExists)
		return &resp, nil
	}

	resp.Msg = fmt.Sprintf(s.messages.Responses.AddedSuccessfully, userId)
	return &resp, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, r *api.EmptyParams) (*api.UsersResponse, error) {
	users, err := s.Service.GetUsers(ctx)
	if err != nil {
		return &api.UsersResponse{}, nil
	}

	return users, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, ue *api.UserEmail) (*api.UserResponse, error) {
	var resp api.UserResponse

	err := ue.Validate()
	if err != nil {
		resp.Msg = s.getMessageError(errInvalidEmail)
		return &resp, nil
	}

	userId, err := s.Service.DeleteUser(ctx, ue)
	if err != nil {
		resp.Msg = s.getMessageError(err)
		return &resp, nil
	}

	if userId == 0 {
		resp.Msg = s.getMessageError(errNotExists)
		return &resp, nil
	}

	resp.Msg = fmt.Sprintf(s.messages.Responses.RemovedSuccessfully, userId)
	return &resp, nil
}
