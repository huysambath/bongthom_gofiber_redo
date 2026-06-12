package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type AuthRoute struct {
	Handler *AuthHandler
}

func NewAuthRoute(app *fiber.App, db *sqlx.DB) *AuthRoute{
	s := NewAuthServiceImpl(db)
	h := NewAuthHandler(s)
	app.Post("/api/v1/admin/auth/login",h.Login)
	return &AuthRoute{
		Handler: h,
	}
}