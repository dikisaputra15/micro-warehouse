package usecase

import (
	"context"
	"fmt"
	"micro-warehouse/transaction-service/model"
	"micro-warehouse/transaction-service/pkg/httpclient"
	"micro-warehouse/transaction-service/pkg/rabbitmq"
	"micro-warehouse/transaction-service/repository"

	"github.com/gofiber/fiber/v2/log"
)

type TransactionUsecaseInterface interface {
	GetDashboardStats(ctx context.Context) (int64, int64, int64, error)
	GetDashboardStatsByMerchant(ctx context.Context, merchantID uint) (int64, int64, int64, error)

	GetTransactions(ctx context.Context, page, limit int, search, sortBy, sortOrder string, merchantID uint) ([]model.Transaction, int64, error)
	CreateTransaction(ctx context.Context, transaction model.Transaction) (int64, error)

	// Midtrans update status transaction
	UpdatePaymentStatus(ctx context.Context, orderID string, paymentStatus, paymentMethod, transactionId, fraudStatus string) error
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepositoryInterface
	merchantClient  *httpclient.MerchantClient
	rabbitMQService *rabbitmq.RabbitMQService
	productClient   *httpclient.ProductClient
	userClient      *httpclient.UserClient
}

// CreateTransaction implements TransactionUsecaseInterface.
func (t *transactionUsecase) CreateTransaction(ctx context.Context, transaction model.Transaction) (int64, error) {
	panic("unimplemented")
}

// GetDashboardStats implements TransactionUsecaseInterface.
func (t *transactionUsecase) GetDashboardStats(ctx context.Context) (int64, int64, int64, error) {
	panic("unimplemented")
}

// GetDashboardStatsByMerchant implements TransactionUsecaseInterface.
func (t *transactionUsecase) GetDashboardStatsByMerchant(ctx context.Context, merchantID uint) (int64, int64, int64, error) {
	panic("unimplemented")
}

// GetTransactions implements TransactionUsecaseInterface.
func (t *transactionUsecase) GetTransactions(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string, merchantID uint) ([]model.Transaction, int64, error) {
	panic("unimplemented")
}

// UpdatePaymentStatus implements TransactionUsecaseInterface.
func (t *transactionUsecase) UpdatePaymentStatus(ctx context.Context, orderID string, paymentStatus string, paymentMethod string, transactionId string, fraudStatus string) error {
	panic("unimplemented")
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepositoryInterface, merchantClient *httpclient.MerchantClient, rabbitMQService *rabbitmq.RabbitMQService, productClient *httpclient.ProductClient, userClient *httpclient.UserClient) TransactionUsecaseInterface {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		merchantClient:  merchantClient,
		rabbitMQService: rabbitMQService,
		productClient:   productClient,
		userClient:      userClient,
	}
}

func (tu *transactionUsecase) validateProductStocks(ctx context.Context, transaction model.Transaction) error {

	for _, product := range transaction.TransactionProducts {
		
		merchantProduct, err := tu.merchantClient.GetMerchantProductStock(
			ctx,
			transaction.MerchantID,
			product.ProductID,
		)

		if err != nil {
			log.Errorf("[TransactionUsecase] validateProductStocks - API error: %v", err)
			return err
		}

		if merchantProduct == nil {
			return fmt.Errorf("product tidak ditemukan")
		}

		if merchantProduct.Stock < int(product.Quantity) {
			log.Warnf(
				"Stock tidak cukup untuk product %d. required: %d, available: %d",
				product.ProductID,
				product.Quantity,
				merchantProduct.Stock,
			)

			return fmt.Errorf(
				"stock tidak mencukupi untuk product '%s'. Dibutuhkan: %d, Tersedia: %d",
				merchantProduct.ProductName,
				product.Quantity,
				merchantProduct.Stock,
			)
		}

		log.Infof("Stock aman untuk product %d", product.ProductID)
	}

	return nil
}
