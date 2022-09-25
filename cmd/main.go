package main

import (
	"assignment2/config"
	"assignment2/httpserver"
	"assignment2/httpserver/controllers"
	"assignment2/httpserver/repositories/gorm"
	"assignment2/httpserver/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}
	orderRepo := gorm.NewOrderRepo(db)
	orderSvc := services.NewOrderSvc(orderRepo)
	orderHandler := controllers.NewOrderController(orderSvc)

	router := gin.Default()
	app := httpserver.NewRouter(router, orderHandler)
	app.Start(":4000")
}
