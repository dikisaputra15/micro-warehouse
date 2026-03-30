package controller

import (
	"micro-warehouse/warehouse-service/controller/request"
	"micro-warehouse/warehouse-service/controller/response"
	"micro-warehouse/warehouse-service/model"
	"micro-warehouse/warehouse-service/pkg/conv"
	"micro-warehouse/warehouse-service/pkg/httpclient"
	"micro-warehouse/warehouse-service/pkg/validator"
	"micro-warehouse/warehouse-service/usecase"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/log"
)

type WarehouseProductControllerInterface interface {
	GetDetailWarehouse(c *fiber.Ctx) error
	GetDetailWarehouseProductByID(c *fiber.Ctx) error
	CreateWarehouseProduct(c *fiber.Ctx) error
	GetWarehouseProductByWarehouseIDAndProductID(c *fiber.Ctx) error
	UpdateWarehouseProduct(c *fiber.Ctx) error
	DeleteWarehouseProduct(c *fiber.Ctx) error
	DeleteAllWarehouseProductByProductID(c *fiber.Ctx) error
	GetWarehouseProductByProductID(c *fiber.Ctx) error
	GetProductTotalStock(c *fiber.Ctx) error
}

type warehouseProductController struct {
	warehouseProductUsecase usecase.WarehouseProductUsecaseInterface
}

// CreateWarehouseProduct implements WarehouseProductControllerInterface.
func (w *warehouseProductController) CreateWarehouseProduct(c *fiber.Ctx) error {
	ctx := c.Context()

	var req request.CreateWarehouseProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("[WarehouseProductController] CreateWarehouseProduct -1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[WarehouseProductController] CreateWarehouseProduct -2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	warehouseID := c.Params("warehouse_id")
	warehouseIDUint := conv.StringToUint(warehouseID)

	reqModel := model.WarehouseProduct{
		WarehouseID: warehouseIDUint,
		ProductID: req.ProductID,
		Stock: req.Stock,
	}

	if err := w.warehouseProductUsecase.CreateWarehouseProduct(ctx, &reqModel); err != nil {
		log.Errorf("[WarehouseProductController] CreateWarehouseProduct -3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create warehouse product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "warehouse product created successfully",
	})
}

// DeleteAllWarehouseProductByProductID implements WarehouseProductControllerInterface.
func (w *warehouseProductController) DeleteAllWarehouseProductByProductID(c *fiber.Ctx) error {
	ctx := c.Context()
	productID := c.Params("product_id")
	productIDUint := conv.StringToUint(productID)

	if err := w.warehouseProductUsecase.DeleteAllWarehouseProductByProductID(ctx, productIDUint); err != nil {
		log.Errorf("[WarehouseProductController] DeleteAllWarehouseProductByProductID - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete all warehouse product by product id",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All warehouse product deleted successfully",
	})
}

// DeleteWarehouseProduct implements WarehouseProductControllerInterface.
func (w *warehouseProductController) DeleteWarehouseProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	warehouseProductID := c.Params("warehouse_product_id")
	warehouseProductIDUint := conv.StringToUint(warehouseProductID)

	if err := w.warehouseProductUsecase.DeleteWarehouseProduct(ctx, warehouseProductIDUint); err != nil {
		log.Errorf("[WarehouseProductController] DeleteWarehouseProduct - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete warehouse product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse product deleted successfully",
	})
}

// GetDetailWarehouse implements WarehouseProductControllerInterface.
func (w *warehouseProductController) GetDetailWarehouse(c *fiber.Ctx) error {
	ctx := c.Context()
	warehouseID := c.Params("warehouse_id")
	warehouseIDUint := conv.StringToUint(warehouseID)

	warehouse, products, err := w.warehouseProductUsecase.GetDetailWarehouse(ctx, warehouseIDUint)
	if err != nil {
		log.Errorf("[WarehouseProductController] GetDetailWarehouse - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get detail warehouse",
		})
	}

	respWarehouseProducts := response.DetailWarehouseResponse{
		ID: warehouse.ID,
		Name: warehouse.Name,
		Address: warehouse.Address,
		Photo: warehouse.Photo,
		Phone: warehouse.Phone,
	}

	productMap := make(map[uint]*httpclient.ProductResponse)
	for i := range products {
		productMap[products[i].ID] = &products[i]
	}

	for _, wp := range warehouse.WarehouseProducts {
		warehouseProduct := response.WarehouseProductResponse{
			ID: wp.ID,
			WarehouseID: wp.WarehouseID,
			ProductID: wp.ProductID,
			Stock: wp.Stock,
		}

		if product, exist := productMap[wp.ProductID]; exist {
			warehouseProduct.ProductName = product.Name
			warehouseProduct.ProductAbout = product.About
			warehouseProduct.ProductPhoto = product.Thumbnail
			warehouseProduct.ProductPrice = int(product.Price)
			warehouseProduct.ProductCategory = product.Category.Name
			warehouseProduct.ProductCategoryPhoto = product.Category.Photo
		}

		respWarehouseProducts.WarehouseProducts = append(respWarehouseProducts.WarehouseProducts, warehouseProduct)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": respWarehouseProducts,
		"message": "Warehouse products fetched successfully",
	})
}

// GetDetailWarehouseProductByID implements WarehouseProductControllerInterface.
func (w *warehouseProductController) GetDetailWarehouseProductByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetProductTotalStock implements WarehouseProductControllerInterface.
func (w *warehouseProductController) GetProductTotalStock(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetWarehouseProductByProductID implements WarehouseProductControllerInterface.
func (w *warehouseProductController) GetWarehouseProductByProductID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetWarehouseProductByWarehouseIDAndProductID implements WarehouseProductControllerInterface.
func (w *warehouseProductController) GetWarehouseProductByWarehouseIDAndProductID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateWarehouseProduct implements WarehouseProductControllerInterface.
func (w *warehouseProductController) UpdateWarehouseProduct(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewWarehouseProductController(warehouseProductUsecase usecase.WarehouseProductUsecaseInterface) WarehouseProductControllerInterface {
	return &warehouseProductController{
		warehouseProductUsecase: warehouseProductUsecase,
	}
}
