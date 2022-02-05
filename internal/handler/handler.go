package handler

import (
	"github.com/s3rzh/go-grpc-user-service/internal/handler/grpc"
	"github.com/s3rzh/go-grpc-user-service/internal/service"
)

type Handler struct {
	Server *grpc.UserManagementServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Server: grpc.NewUserManagementServer(service)}
}
