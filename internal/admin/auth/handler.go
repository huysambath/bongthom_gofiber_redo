package auth

import (
	"admin-api/pkg/utls"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(db *sqlx.DB) *AuthHandler{
	s := NewAuthServiceImpl(db)
	return &AuthHandler{
		Service: s,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req AuthRequest
	v := utls.NewValidator()

	if err := req.bind(c,v); err != nil {
		return err
	}

	result, err := h.Service.Login(req.Username,req.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "failed",
			"message": "error",
			"status": 2000,
			"data": "JWT not generated",
		})
	}

	if result == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": "fail",
			"message": "success",
			"status": 3001,
			"data": result,
		})
	}else{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": "true",
			"message": "success",
			"status": 3000,
			"data": result,
		})
	}
} 