package controller

import (
	"micro-warehouse/merchant-service/controller/response"
	"micro-warehouse/merchant-service/pkg/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UploadControllerInterface interface {
	UploadMerchantPhoto(c *fiber.Ctx) error
}

type UploadController struct {
	fileUploadHelper *storage.FileUploadHelper
}

// UploadMerchantPhoto implements UploadControllerInterface.
func (u *UploadController) UploadMerchantPhoto(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		log.Errorf("Failed to get file %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get file",
			"error": err.Error(),
		})
	}

	result, err := u.fileUploadHelper.UploadPhoto(c.Context(), file, "merchants")
	if err != nil {
		log.Errorf("Failed to upload file: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload file",
			"error": err.Error(),
		})
	}

	response := response.UploadResponse{
		URL: result.URL,
		Path: result.Path,
		Filename: result.Filename,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"data": response,
	})
}

func NewUploadController(fileUploadHelper *storage.FileUploadHelper) UploadControllerInterface {
	return &UploadController{
		fileUploadHelper: fileUploadHelper,
	}
}
