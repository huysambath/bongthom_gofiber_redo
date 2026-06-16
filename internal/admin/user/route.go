package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type UserRoute struct {
	Handler *UserHandler
}

func NewUserRoute(app *fiber.App, db *sqlx.DB) *UserRoute {
	repo := NewUserRepoImpl(db)
	svc := NewUserServiceImpl(repo)
	h := NewUserHandler(svc)

	v1 := app.Group("/api/v1/admin")
	users := v1.Group("/users")
	users.Get("/", h.Show)
	users.Get("/:id",h.ShowOne)
	users.Post("/create",h.Create)
	users.Put("/update/:id",h.Update)
	users.Delete("/delete/:id",h.Delete)
	return &UserRoute {
		Handler: h,
	}
}