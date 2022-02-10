package queue

import (
	"github.com/s3rzh/go-grpc-user-service/pkg/queue/rabbitmq"
	"github.com/streadway/amqp"
)

type Queue interface {
	Send([]byte) error
	Receive() (<-chan amqp.Delivery, error)
	Close() error
}

func NewQueue(cfg rabbitmq.Config) (Queue, error) {
	return rabbitmq.NewRabbitmq(cfg)
}
