package controller

import (
	"micro-warehouse/warehouse-service/controller/response"
	"micro-warehouse/warehouse-service/pkg/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UploadControllerInterface interface {
	UploadPhoto(ctx *fiber.Ctx) error
}

type uploadController struct {
	fileUploadHelper *storage.FileUploadHelper
}

// UploadPhoto implements UploadControllerInterface.
func (u *uploadController) UploadPhoto(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		log.Errorf("Failed to get file %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get file",
			"error": err.Error(),
		})
	}

	result, err := u.fileUploadHelper.UploadPhoto(ctx.Context(), file, "warehouses")
	if err != nil {
		log.Errorf("Failed to upload file: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload file",
			"error": err.Error(),
		})
	}

	response := response.UploadResponse{
		URL: result.URL,
		Path: result.Path,
		Filename: result.Filename,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"data": response,
	})
}

func NewUploadController(fileUploadHelper *storage.FileUploadHelper) UploadControllerInterface {
	return &uploadController{
		fileUploadHelper: fileUploadHelper,
	}
}
