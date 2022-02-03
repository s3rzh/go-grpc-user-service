package app

import (
	"fmt"
	"log"

	"github.com/s3rzh/go-grpc-user-service/internal/config"
	"github.com/s3rzh/go-grpc-user-service/internal/server"
)

func Run(configPath string) {
	cfg, err := config.InitApp(configPath)
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//log.Printf(" hello from app!%+v", cfg)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Port); err != nil {
			log.Fatalf("error running server: %s", err.Error())
		}
	}()

	// if err := srv.Run(cfg.Port); err != nil {
	// 	log.Fatalf("error running server: %s", err.Error())
	// }

	cn := make(chan struct{}, 1)
	<-cn

	fmt.Println("App end!")
}
