package handler

import (
	"admin-api/internal/admin/auth"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ServiceHandlers struct {
	Admin *AdminService
	Front *FrontService
}

func NewServiceHandlers(a *fiber.App, db *sqlx.DB, rdb *redis.Client) *ServiceHandlers{
	return &ServiceHandlers{
		Admin: NewAdminService(a,db,rdb),
		Front: NewFrontService(a),
	}
}

type AdminService struct {
	Auth *auth.AuthRoute
}

func NewAdminService(a *fiber.App, db *sqlx.DB, rdb *redis.Client) *AdminService{
	authRoute := auth.NewAuthRoute(a,db,rdb)
	return &AdminService{
		Auth: authRoute,
	}
}

type FrontService struct {

}

func NewFrontService(a *fiber.App) *FrontService{
	return &FrontService{

	}
}