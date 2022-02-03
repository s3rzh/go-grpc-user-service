package app

import (
	"fmt"
	"log"

	"github.com/s3rzh/go-grpc-user-service/internal/config"
	"github.com/s3rzh/go-grpc-user-service/internal/handler"
	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/internal/server"
	"github.com/s3rzh/go-grpc-user-service/internal/service"
)

func Run(configPath string) {
	cfg, err := config.InitApp(configPath)
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//log.Printf(" hello from app!%+v", cfg)

	repository := repository.NewRepository()
	service := service.NewService(repository)
	//userServer := grpc.NewUserManagementServer(service)

	handler := handler.NewHandler(service)
	//userServer1.Service

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Port, handler.Service); err != nil {
			log.Fatalf("error running server: %s", err.Error())
		}
	}()

	// if err := srv.Run(cfg.Port); err != nil {
	// 	log.Fatalf("error running server: %s", err.Error())
	// }

	cn := make(chan struct{}, 1)
	<-cn

	fmt.Println("app ended..")
}
