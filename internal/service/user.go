package service

import (
	"context"

	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserGRPCService struct {
	rep *repository.Repository
}

func NewUserGRPCService(r *repository.Repository) *UserGRPCService {
	return &UserGRPCService{rep: r}
}

func (s *UserGRPCService) CreateUser(ctx context.Context, u *api.User) (int, error) {
	// check for cache
	var userId int

	userId, err := s.rep.GetUserIdByEmail(ctx, u.Email)
	if err != nil {
		return 0, err
	}

	if userId > 0 {
		return 0, nil
	}

	userId, err = s.rep.CreateUser(ctx, u)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *UserGRPCService) GetUsers(ctx context.Context) (*api.UsersResponse, error) {
	users, err := s.rep.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserGRPCService) DeleteUser(ctx context.Context, ue *api.UserEmail) (int, error) {

	userId, err := s.rep.GetUserIdByEmail(ctx, ue.Email)
	if err != nil {
		return 0, err
	}

	if userId == 0 {
		return 0, nil
	}

	err = s.rep.DeleteUser(ctx, ue.Email)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
