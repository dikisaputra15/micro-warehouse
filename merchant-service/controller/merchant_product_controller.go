package controller

import (
	"micro-warehouse/merchant-service/controller/request"
	"micro-warehouse/merchant-service/model"
	"micro-warehouse/merchant-service/pkg/conv"
	"micro-warehouse/merchant-service/pkg/httpclient"
	"micro-warehouse/merchant-service/pkg/validator"
	"micro-warehouse/merchant-service/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type MerchantProductControllerInterface interface {
	CreateMerchantProduct(c *fiber.Ctx) error
	GetMerchanProductByID(c *fiber.Ctx) error
	GetMerchantProducts(c *fiber.Ctx) error
	GetMerchantProductByBarcode(c *fiber.Ctx) error
	UpdateMerchantProduct(c *fiber.Ctx) error
	DeleteMerchantProduct(c *fiber.Ctx) error
	DeleteAllProductMerchantProducts(c *fiber.Ctx) error
	GetProductTotalStock(c *fiber.Ctx) error
}

type merchantProductController struct {
	merchantProductUsecase usecase.MerchantProductUsecaseInterface
}

// CreateMerchantProduct implements MerchantProductControllerInterface.
func (m *merchantProductController) CreateMerchantProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	var req request.CreateMerchantProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("[MerchantProductController] CreateMerchantProduct - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[MerchantProductController] CreateMerchantProduct - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.MerchantProduct{
		ProductID:  req.ProductID,
		WarehouseID: req.WarehouseID,
		Stock:       req.Stock,
		MerchantID:  req.MerchantID,
	}

	if err := m.merchantProductUsecase.CreateMerchantProduct(ctx, &reqModel); err != nil {
		log.Errorf("[MerchantProductController] CreateMerchantProduct - 3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create merchant product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Merchant product created successfully",
	})
}

// DeleteAllProductMerchantProducts implements MerchantProductControllerInterface.
func (m *merchantProductController) DeleteAllProductMerchantProducts(c *fiber.Ctx) error {
	ctx := c.Context()
	productID := c.Params("product_id")
	productIDUint := conv.StringToUint(productID)

	if err := m.merchantProductUsecase.DeleteAllProductMerchantProducts(ctx, productIDUint); err != nil {
		log.Errorf("[MerchantProductController] DeleteAllProductMerchantProducts - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete all product merchant products",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All product merchant products deleted successfully",
	})

}

// DeleteMerchantProduct implements MerchantProductControllerInterface.
func (m *merchantProductController) DeleteMerchantProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	merchantProductID := c.Params("merchant_product_id")
	merchantProductIDUint := conv.StringToUint(merchantProductID)

	if err := m.merchantProductUsecase.DeleteMerchantProduct(ctx, merchantProductIDUint); err != nil {
		log.Errorf("[MerchantProductController] DeleteMerchantProduct - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete merchant product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Merchant product deleted successfully",
	})
}

// GetMerchanProductByID implements MerchantProductControllerInterface.
func (m *merchantProductController) GetMerchanProductByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetMerchantProductByBarcode implements MerchantProductControllerInterface.
func (m *merchantProductController) GetMerchantProductByBarcode(c *fiber.Ctx) error {
	ctx := c.Context()
	barcode := c.Params("barcode")
	merchantID := c.Params("merchant_id")
	merchantIDUint := conv.StringToUint(merchantID)
	if merchantIDUint == 0 {
		merchantIDUint = conv.StringToUint(c.Query("merchant_id"))
	}

	merchantProduct, product, warehouse, err := m.merchantProductUsecase.GetMerchantProductByBarcode(ctx, barcode, merchantIDUint)
	if err != nil {
		log.Errorf("[MerchantProductController] GetMerchantProductByBarcode - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get merchant product by barcode",
		})
	}

	productResponse := httpclient.MapProductResponseToMerchantProduct(product)
	warehouseResponse := httpclient.MapWarehouseResponseToMerchantProduct(warehouse)

	productResponse.ID = merchantProduct.ID
	productResponse.MerchantID = merchantProduct.MerchantID
	productResponse.ProductID = merchantProduct.ProductID
	productResponse.Stock = merchantProduct.Stock
	productResponse.WarehouseID = merchantProduct.WarehouseID
	productResponse.WarehouseName = warehouseResponse.WarehouseName
	productResponse.WarehousePhoto = warehouseResponse.WarehousePhoto
	productResponse.WarehousePhone = warehouseResponse.WarehousePhone

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Merchant product retrieved successfully",
		"data":    productResponse,
	})
}

// GetMerchantProducts implements MerchantProductControllerInterface.
func (m *merchantProductController) GetMerchantProducts(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetProductTotalStock implements MerchantProductControllerInterface.
func (m *merchantProductController) GetProductTotalStock(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateMerchantProduct implements MerchantProductControllerInterface.
func (m *merchantProductController) UpdateMerchantProduct(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewMerchantProductController(merchantProductUsecase usecase.MerchantProductUsecaseInterface) MerchantProductControllerInterface {
	return &merchantProductController{
		merchantProductUsecase: merchantProductUsecase,
	}
}
