package rabbitmq

import (
	"fmt"
	"micro-warehouse/warehouse-service/repository"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn *amqp.Connection
	channel *amqp.Channel
	repo repository.WarehouseProductRepositoryInterface
}

type StockReductionEvent struct {
	WarehouseID uint `json:"warehouse_id"`
	ProductID uint `json:"product_id"`
	Stock uint `json:"stock"`
	MerchantID uint `json:"merchant_id"`
	Timestamp uint `json:"timestamp"`
}

const (
	ExchangeName = "warehouse_events"
	QueueName = "stock_reduction_queue"
	RoutingKey = "stock.reduction"
)

func NewRabbitMQConsumer(rabbitMQURL string, repo repository.WarehouseProductRepositoryInterface) (*RabbitMQConsumer, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare exchange
	err = ch.ExchangeDeclare(
		ExchangeName, // name
		"topic", // type
		true, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare queue
	q, err := ch.QueueDeclare(
		QueueName, // name
		true, // durable
		false, // delete when unused
		false, //exclusive
		false, // no-wait
		nil, // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind queue to exchange
	err = ch.QueueBind(
		q.Name, // queue name
		RoutingKey, // routing key
		ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &RabbitMQConsumer{
		conn: conn,
		channel: ch,
		repo: repo,
	}, nil
}