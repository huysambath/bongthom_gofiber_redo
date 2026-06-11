package handler

import (
	"admin-api/internal/admin/auth"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type ServiceHandlers struct {
	Admin *AdminService
	Front *FrontService
}

func NewServiceHandlers(a *fiber.App, db *sqlx.DB) *ServiceHandlers{
	return &ServiceHandlers{
		Admin: NewAdminService(a,db),
		Front: NewFrontService(a),
	}
}

type AdminService struct {
	Auth *auth.AuthRoute
}

func NewAdminService(a *fiber.App, db *sqlx.DB) *AdminService{
	authRoute := auth.NewAuthRoute(a,db)
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