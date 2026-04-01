package request

type CreateWarehouseProductRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Stock     int  `json:"stock" validate:"required"`
}

type UpdateWarehouseProductRequest struct {
	ProductID *uint `json:"product_id,omitempty"`
	Stock     *int  `json:"stock,omitempty"`
}