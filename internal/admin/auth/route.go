package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthRoute struct {
	Handler *AuthHandler
}

func NewAuthRoute(app *fiber.App, db *sqlx.DB, rdb *redis.Client) *AuthRoute{
	s := NewAuthServiceImpl(db, rdb)
	h := NewAuthHandler(s)
	app.Post("/api/v1/admin/auth/login",h.Login)
	return &AuthRoute{
		Handler: h,
	}
}