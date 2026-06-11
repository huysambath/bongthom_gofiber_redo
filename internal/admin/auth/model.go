package auth

import (
	"admin-api/pkg/utls"

	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"user_name"`
	Password string `json:"password" db:"password"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required"`
}

func (r *AuthRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(&r); err != nil {
		return  err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type AuthResponse struct {
	IsSuccess bool `json:"is_success"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
	TokenType string `json:"token_type"`
}