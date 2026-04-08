package rabbitmq

import (
	"encoding/json"
	"micro-warehouse/merchant-service/repository"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/streadway/amqp"
)

type StockReducedEvent struct {
	MerchantID uint                         `json:"merchant_id"`
	Products   []StockReductionEventProduct `json:"products"`
	OrderID    uint                         `json:"order_id"`
	TimeStamp  time.Time                    `json:"timestamp"`
}

type StockReductionEventProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type StockConsumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	merchantRepo repository.MerchantRepositoryInterface
}

func NewStockConsumer(url string, merchantRepo repository.MerchantRepositoryInterface) *StockConsumer {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Errorf("[StockConsumer] NewStockConsumer - 1: %v", err)
		return nil
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Errorf("[StockConsumer] NewStockConsumer - 2: %v", err)
		return nil
	}

	err = ch.ExchangeDeclare(
		"business_events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[StockConsumer] NewStockConsumer - 3: %v", err)
		return nil
	}

	q, err := ch.QueueDeclare(
		"merchant_stock_events",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[StockConsumer] NewStockConsumer - 4: %v", err)
		return nil
	}

	err = ch.QueueBind(
		q.Name,
		"merchant.stock.*",
		"business_events",
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[StockConsumer] NewStockConsumer - 5: %v", err)
		return nil
	}

	return &StockConsumer{
		conn: conn,
		ch:   ch,
		merchantRepo: merchantRepo,
	}
}

func (s *StockConsumer) ConsumereStockReductionEvents(ctx context.Context) error {
	msgs, err := s.ch.Consume(
		"merchant_stock_events",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[StockConsumer] ConsumereStockReductionEvents - 1: %v", err)
		return err
	}

	for {
		select {
			case <-ctx.Done():
				log.Info("Stoping stock consumer ...")
				return nil
			case msg := <-msgs:
				go sc.handle
		}
	}

}

func (sc *StockConsumer) handleStockReductionEvent(msg amqp.Delivery) error {
	defer msg.Ack(false)

	var event StockReducedEvent
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		log.Errorf("[StockConsumer] handleStockReductionEvent - 1: %v", err)
		return err
	}
} 