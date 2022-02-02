package server

type Server struct {
}

func (s *Server) Run(port string) error {

	//listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	// if err != nil {
	// 	return err
	// }

	//grpcServer := grpc.NewServer()
	//srv := &user

	//api.RegisterUserManagementServer()

	println("server started!")
	return nil
}
