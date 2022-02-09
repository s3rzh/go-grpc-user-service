package queue

import (
	"github.com/s3rzh/go-grpc-user-service/pkg/queue/rabbitmq"
)

type Queue interface {
	Send(string) error
	Close() error
}

func NewQueue(cfg rabbitmq.Config) (Queue, error) {
	return rabbitmq.NewRabbitmq(cfg)
}
