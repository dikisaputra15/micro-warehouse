package repository

import (
	"context"
	"micro-warehouse/merchant-service/model"
)

type MerchantProductRepositoryInterface interface {
	CreateMerchantProduct(ctx context.Context, merchantProduct *model.MerchantProduct) error
	GetMerchantProductByID(ctx context.Context, id uint) (*model.MerchantProduct, error)
	GetMerchantProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string, merchantID, productID uint) ([]*model.MerchantProduct, int64, error)
	GetMerchantProductsByProductIDAndMerchantID(ctx context.Context, productID uint, merchantID uint) (*model.MerchantProduct, error)
	UpdateMerchantProduct(ctx context.Context, merchantProduct *model.MerchantProduct) error
	DeleteMerchantProduct(ctx context.Context, id uint) error
	DeleteAllProductMerchantProducts(ctx context.Context, productID uint) error

	GetProductTotalStock(ctx context.Context, productID uint) (int, error)
	ReduceStock(ctx context.Context, merchantID uint, productID uint, quantity int64) error
}