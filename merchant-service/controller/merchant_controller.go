package controller

import (
	"micro-warehouse/merchant-service/usecase"

	"github.com/gofiber/fiber/v2"
)

type MerchantControllerInerface interface {
	CreateMerchant(c *fiber.Ctx) error
	GetAllMerchants(c *fiber.Ctx) error
	GetMerchantByID(c *fiber.Ctx) error
	UpdateMerchant(c *fiber.Ctx) error
	DeleteMerchant(c *fiber.Ctx) error
}

type merchantController struct {
	merchantUsecase usecase.MerchantUsecaseInterface
}

// CreateMerchant implements MerchantControllerInerface.
func (m *merchantController) CreateMerchant(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteMerchant implements MerchantControllerInerface.
func (m *merchantController) DeleteMerchant(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllMerchants implements MerchantControllerInerface.
func (m *merchantController) GetAllMerchants(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetMerchantByID implements MerchantControllerInerface.
func (m *merchantController) GetMerchantByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateMerchant implements MerchantControllerInerface.
func (m *merchantController) UpdateMerchant(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewMerchantController(merchantUsecase usecase.MerchantUsecaseInterface) MerchantControllerInerface {
	return &merchantController{
		merchantUsecase: merchantUsecase,
	}
}
