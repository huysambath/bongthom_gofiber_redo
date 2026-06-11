package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func New() *fiber.App {
	f := fiber.New(fiber.Config{

	})

	f.Use(logger.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: []string {"*"},
		AllowHeaders: []string {"Origins", "Content-Type", "Accept", "Authorization", "Accept-Language"},
		AllowMethods: []string {"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
	}))
	return f
}