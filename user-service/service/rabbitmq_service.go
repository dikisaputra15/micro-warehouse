package service

import (
	"context"
	"encoding/json"
	"fmt"
	"micro-warehouse/user-service/configs"

	"github.com/gofiber/fiber/v2/log"
	"github.com/streadway/amqp"
)

type EmailPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
}

type RabbitMQServiceInterface interface {
	PublishEmail(ctx context.Context, payload EmailPayload) error
	Close() error
}

type rabbitMQService struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	config configs.Config
}

// Close implements RabbitMQServiceInterface.
func (r *rabbitMQService) Close() error {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}

// PublishEmail implements RabbitMQServiceInterface.
func (r *rabbitMQService) PublishEmail(ctx context.Context, payload EmailPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal email payload: %v", err)
	}

	queue, err := r.ch.QueueDeclare(
		"email_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare email queue: %v", err)
	}

	err = r.ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish email message: %v", err)
	}
	return nil
}

func NewRabbitMQService(config configs.Config) (RabbitMQServiceInterface, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQ.Username, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port))
	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 1: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 2: %v", err)
		return nil, err
	}
	return &rabbitMQService{
		conn:   conn,
		ch:     ch,
		config: config,
	}, nil
}
