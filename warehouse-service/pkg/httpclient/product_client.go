package httpclient

import (
	"context"
	"micro-warehouse/warehouse-service/configs"
	"net/http"
)

type ProductClientInterface interface {
	GetProductByID(ctx context.Context, productID uint) (*ProductResponse, error)
	GetProductsByIDs(ctx context.Context, productIDs []uint) ([]ProductResponse, error)
	GetProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]ProductResponse, error)
	HealthCheck(ctx context.Context) error
}

type ProductClient struct {
	urlProductService string
	httpClient *http.Client
}

// GetProductByID implements ProductClientInterface.
func (p *ProductClient) GetProductByID(ctx context.Context, productID uint) (*ProductResponse, error) {
	panic("unimplemented")
}

// GetProducts implements ProductClientInterface.
func (p *ProductClient) GetProducts(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]ProductResponse, error) {
	panic("unimplemented")
}

// GetProductsByIDs implements ProductClientInterface.
func (p *ProductClient) GetProductsByIDs(ctx context.Context, productIDs []uint) ([]ProductResponse, error) {
	panic("unimplemented")
}

// HealthCheck implements ProductClientInterface.
func (p *ProductClient) HealthCheck(ctx context.Context) error {
	panic("unimplemented")
}

type ProductResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	About     string `json:"about"`
	Price     int64  `json:"price"`
	Thumbnail string `json:"thumbnail"`
	Category  struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	} `json:"category"`
}

type ProductServiceResponse struct {
	Message string          `json:"message"`
	Data    ProductResponse `json:"data"`
	Error   string          `json:"error,omitempty"`
}

type ProductListResponse struct {
	Message string            `json:"message"`
	Data    []ProductResponse `json:"data"`
	Error   string            `json:"error,omitempty"`
}

func NewProductClient(httpClient *http.Client, cfg configs.Config) ProductClientInterface {
	return &ProductClient{httpClient: httpClient}
}
