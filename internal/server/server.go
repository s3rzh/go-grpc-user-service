package server

import (
	"fmt"
	"net"

	usrv "github.com/s3rzh/go-grpc-user-service/internal/handler/grpc"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
}

func (s *Server) Run(port string) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	s.listener = ln

	grpcServer := grpc.NewServer()
	userServer := &usrv.UserManagementServer{}

	api.RegisterUserManagementServer(grpcServer, userServer)

	err = grpcServer.Serve(s.listener)
	if err != nil {
		return err
	}

	fmt.Println("Server started2!")
	return nil
}

func (s *Server) Stop() {
	s.listener.Close()
}
