package app

import (
	"micro-warehouse/merchant-service/configs"
	"micro-warehouse/merchant-service/controller"
	"micro-warehouse/merchant-service/database"
	"micro-warehouse/merchant-service/pkg/redis"
	"micro-warehouse/merchant-service/repository"
	"micro-warehouse/merchant-service/usecase"

	"github.com/gofiber/fiber/v2/log"
)

type Container struct {
	MerchantController        controller.MerchantControllerInerface
	MerchantProductController controller.MerchantProductControllerInterface
}

func BuildContainer() *Container {
	cfg := configs.NewConfig()
	db, err := database.ConnectionPostgres(*cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	redisClient := redis.NewRedisClient(*cfg)
	rabbitMQService, err := service.NewRabbitMQConsumer(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	} 

	cachedUserClient := redis.NewRedisClient(*cfg)

	merchantRepo := repository.NewMerchantRepository(db.DB)
	merchantUsecase := usecase.NewMerchantUsecase(merchantRepo)
	merchantController := controller.NewMerchantController(merchantUsecase)
}