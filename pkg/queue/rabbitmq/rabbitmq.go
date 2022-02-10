package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Config struct {
	Host      string
	Port      string
	Username  string
	Password  string
	QueueName string
}

type Rabbitmq struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbitmq(cfg Config) (*Rabbitmq, error) {
	var url string = fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		cfg.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Rabbitmq{conn: conn, ch: ch, q: q}, nil

}

func (r *Rabbitmq) Send(msg []byte) error {
	err := r.ch.Publish(
		"",
		r.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Rabbitmq) Receive() (<-chan amqp.Delivery, error) {
	msgs, err := r.ch.Consume(
		r.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *Rabbitmq) Close() error {
	err := r.ch.Close()
	if err != nil {
		return err
	}

	err = r.conn.Close()
	if err != nil {
		return err
	}

	return nil
}
