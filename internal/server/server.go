package server

import (
	"fmt"
	"net"

	usr "github.com/s3rzh/go-grpc-user-service/internal/handler/grpc"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
}

func (s *Server) Run(port string, userServer *usr.UserManagementServer) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	s.listener = ln

	grpcServer := grpc.NewServer()
	//userServer := usr.NewUserManagementServer()

	api.RegisterUserManagementServer(grpcServer, userServer)

	err = grpcServer.Serve(s.listener)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	err := s.listener.Close()
	if err != nil {
		return err
	}
	return nil
}
