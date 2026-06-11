package main

import (
	"admin-api/config"
	"admin-api/config/database"
	"admin-api/handler"
	"admin-api/router"
	"fmt"
)

func main(){
	app_configs := config.NewConfig()
	app := router.New()
	db_pool := database.GetDB()
	

	handler.NewServiceHandlers(app,db_pool)
	app.Listen(fmt.Sprintf("%s:%d", app_configs.AppHost, app_configs.AppPort))
}