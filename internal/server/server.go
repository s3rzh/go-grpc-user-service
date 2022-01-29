package server

import (
	"fmt"
	"net"

	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Run(port string) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	//srv := &user

	api.RegisterUserManagementServer()

	println("server started!")
	return nil
}
