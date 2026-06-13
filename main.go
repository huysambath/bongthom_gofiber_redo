package main

import (
	"admin-api/config"
	"admin-api/config/database"
	"admin-api/config/redis"
	"admin-api/handler"
	"admin-api/router"
	"fmt"
)

func main(){
	app_configs := config.NewConfig()
	app := router.New()
	db_pool := database.GetDB()
	
	rdb := redis.NewRedisClient()

	handler.NewServiceHandlers(app,db_pool,rdb)
	app.Listen(fmt.Sprintf("%s:%d", app_configs.AppHost, app_configs.AppPort))
}