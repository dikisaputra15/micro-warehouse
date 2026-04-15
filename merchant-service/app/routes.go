package app

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, c *Container) {
	api := app.Group("/api/v1")

	merchant := api.Group("/merchants")
	merchant.Post("/", c.MerchantController.CreateMerchant)
	merchant.Get("/", c.MerchantController.GetAllMerchants)
	merchant.Get("/:id", c.MerchantController.GetMerchantByID)
	merchant.Put("/:id", c.MerchantController.UpdateMerchant)
	merchant.Delete("/:id", c.MerchantController.DeleteMerchant)

	merchantProduct := api.Group("/merchant-products")
	merchantProduct.Post("/", c.MerchantProductController.CreateMerchantProduct)
	merchantProduct.Get("/:merchant_product_id", c.MerchantProductController.GetMerchanProductByID)
	merchantProduct.Get("/", c.MerchantProductController.GetMerchantProducts)
	merchantProduct.Get("/barcode/:barcode", c.MerchantProductController.GetMerchantProductByBarcode)
	merchantProduct.Put("/:merchant_product_id", c.MerchantProductController.UpdateMerchantProduct)
	merchantProduct.Delete("/:merchant_product_id", c.MerchantProductController.DeleteMerchantProduct)
	merchantProduct.Delete("/product/:product_id", c.MerchantProductController.DeleteAllProductMerchantProducts)
	merchantProduct.Get("/:product_id/total-stock", c.MerchantProductController.GetProductTotalStock)

	api.Post("/upload-merchant", c.UploadController.UploadMerchantPhoto)
}