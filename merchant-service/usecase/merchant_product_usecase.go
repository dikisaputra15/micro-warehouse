package usecase

import (
	"context"
	"micro-warehouse/merchant-service/model"
	"micro-warehouse/merchant-service/pkg/httpclient"
	"micro-warehouse/merchant-service/repository"
)

type MerchantProductUsecaseInterface interface {
	CreateMerchantProduct(ctx context.Context, merchantProduct *model.MerchantProduct) error
	GetMerchantProductByID(ctx context.Context, merchantProductID uint) (*model.MerchantProduct, error)
	GetMerchantProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string, merchantID, productID uint) ([]model.MerchantProduct, error)
	UpdateMerchantProduct(ctx context.Context, merchantProduct *model.MerchantProduct) error
	DeleteMerchantProduct(ctx context.Context, merchantProductID uint) error
	DeleteAllProductMerchantProducts(ctx context.Context, productID uint) error

	GetProductTotalStock(ctx context.Context, productID uint) (int, error)
}

type merchantProductUsecase struct {
	merchantProductRepo repository.MerchantProductRepositoryInterface
	productClient	   httpclient.ProductClientInterface
	warehouseClient	   httpclient.WarehouseClientInterface
	rabbitMQService *rabbitmq.RabbitMQService
}