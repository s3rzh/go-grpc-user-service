package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/s3rzh/go-grpc-user-service/internal/config"
	"github.com/s3rzh/go-grpc-user-service/internal/handler"
	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/internal/server"
	"github.com/s3rzh/go-grpc-user-service/internal/service"
	"github.com/s3rzh/go-grpc-user-service/pkg/cache"
	"github.com/s3rzh/go-grpc-user-service/pkg/cache/redis"
	"github.com/s3rzh/go-grpc-user-service/pkg/database/postgresql"
	"github.com/s3rzh/go-grpc-user-service/pkg/queue"
	"github.com/s3rzh/go-grpc-user-service/pkg/queue/rabbitmq"
)

func Run(configPath string) {
	cfg, err := config.InitApp(configPath)
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//log.Printf("cfg %+v", cfg)

	db, err := postgresql.NewPostgresDB(postgresql.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	cache, err := cache.NewCache(redis.Config{
		Host:     cfg.Cache.Host,
		Port:     cfg.Cache.Port,
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.DB,
	})
	if err != nil {
		log.Fatalf("failed to initialize cache: %s", err.Error())
	}

	queue, err := queue.NewQueue(rabbitmq.Config{
		Host:      cfg.Queue.Host,
		Port:      cfg.Queue.Port,
		Username:  cfg.Queue.Username,
		Password:  cfg.Queue.Password,
		QueueName: cfg.Queue.QueueName,
	})
	if err != nil {
		log.Fatalf("failed to initialize queue: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository, cache, queue)
	handler := handler.NewHandler(service, cfg.Messages)

	srv := new(server.Server)
	fmt.Println("Started!")
	go func() {
		if err := srv.Run(cfg.Port, handler.Server); err != nil {
			log.Fatalf("error running server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Stop(); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(context.Background()); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}

	//fmt.Println("app ended..")
}
