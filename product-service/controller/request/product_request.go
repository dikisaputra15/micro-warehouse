package request

type CreateProductRequets struct {
	Name string `json:"name" validate:"required"`
	Barcode string `json:"barcode" validate:"required"`
	Price string `json:"price" validate:"required"`
	About string `json:"about" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	Thumbnail string `json:"thumbnail" validate:"required"`
	IsPopular string `json:"is_popular" validate:"required"`
}

type GetAllProductRequest struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Search    string `query:"search"`
	SortBy    string `query:"sort_by"`
	SortOrder string `query:"sort_order"`
}