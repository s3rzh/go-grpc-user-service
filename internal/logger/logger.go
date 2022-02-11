package logger

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/s3rzh/go-grpc-user-service/internal/config"
	"github.com/s3rzh/go-grpc-user-service/internal/entity"
	"github.com/s3rzh/go-grpc-user-service/pkg/database/clickhouse"
	"github.com/s3rzh/go-grpc-user-service/pkg/queue"
	"github.com/s3rzh/go-grpc-user-service/pkg/queue/rabbitmq"
)

const (
	TickerInterval = 10 * time.Second
	BufferSize     = 8
	PartSize       = 4
)

func Run(configPath string) {
	log.Printf("Logger start!")
	cfg, err := config.InitApp(configPath)
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	conn, err := clickhouse.NewClickHouseDB(clickhouse.Config{
		Host:     cfg.ClickHouse.Host,
		Port:     cfg.ClickHouse.Port,
		Username: cfg.ClickHouse.Username,
		DBName:   cfg.ClickHouse.DBName,
		Password: cfg.ClickHouse.Password,
	})
	if err != nil {
		log.Fatalf("failed to initialize clickhouse: %s", err.Error())
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

	msgs, err := queue.Receive()
	if err != nil {
		log.Fatalf("failed to initialize consumer: %s", err.Error())
	}

	buffer := make(chan []byte, BufferSize)

	go func() {
		for d := range msgs {
			buffer <- d.Body
		}
	}()

	ticker := time.NewTicker(TickerInterval)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				l := len(buffer)
				if l > 0 {
					n := PartSize
					if l < n {
						n = l
					}

					batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO users")
					if err != nil {
						log.Fatal(err)
					}

					for i := 0; i < n; i++ {
						var user entity.Data

						err := json.Unmarshal(<-buffer, &user)
						if err != nil {
							log.Fatalf("failed to unmarshal user data: %s", err.Error())
							return
						}

						err = batch.Append(user.ID, user.Email)
						if err != nil {
							log.Fatalf("failed to append user: %s", err.Error())
						}
					}
					if err := batch.Send(); err != nil {
						log.Fatalf("failed to send users: %s", err.Error())
					}
				}
			case <-stop:
				return
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	ticker.Stop()
	stop <- true
	<-stop

	if err := queue.Close(); err != nil {
		log.Fatalf("error occured on queue shutting down: %s", err.Error())
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("error occured on clickhouse shutting down: %s", err.Error())
	}

}
