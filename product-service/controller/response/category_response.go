package response

import "micro-warehouse/product-service/pkg/pagination"

type CategoryRsponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Tagline       string `json:"tagline"`
	Photo       string `json:"photo"`
	CountProduct int  `json:"count_product"`
}

type GetAllCategoriResponse struct {
	Categories  []CategoryRsponse           `json:"categories"`
	Pagination pagination.PaginationResponse `json:"pagination"`
}