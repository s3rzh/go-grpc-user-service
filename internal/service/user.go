package service

import (
	"context"
	"encoding/json"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"github.com/s3rzh/go-grpc-user-service/pkg/cache"
)

const (
	UserList       string        = "user:list"
	ExperationTine time.Duration = 1 * time.Minute
)

type UserGRPCService struct {
	rep   *repository.Repository
	cache cache.Cache
}

func NewUserGRPCService(r *repository.Repository, cache cache.Cache) *UserGRPCService {
	return &UserGRPCService{rep: r, cache: cache}
}

func (s *UserGRPCService) CreateUser(ctx context.Context, u *api.User) (int, error) {
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

	_, err = s.cache.Delete(ctx, UserList)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *UserGRPCService) GetUsers(ctx context.Context) (*api.UsersResponse, error) {
	var users *api.UsersResponse

	cached, err := s.cache.Get(ctx, UserList)
	if err != redis.Nil {
		err := json.Unmarshal([]byte(cached), &users)
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	users, err = s.rep.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	_, err = s.cache.Set(ctx, UserList, bytes, ExperationTine)
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

	_, err = s.cache.Delete(ctx, UserList)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
