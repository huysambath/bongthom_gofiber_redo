package handler

import (
	"admin-api/internal/admin/auth"
	"admin-api/internal/admin/user"

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
	Users *user.UserRoute
}

func NewAdminService(a *fiber.App, db *sqlx.DB, rdb *redis.Client) *AdminService{
	authRoute := auth.NewAuthRoute(a,db,rdb)
	userRoute := user.NewUserRoute(a,db)
	return &AdminService{
		Auth: authRoute,
		Users: userRoute,
	}
}

type FrontService struct {

}

func NewFrontService(a *fiber.App) *FrontService{
	return &FrontService{

	}
}