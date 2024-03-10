package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zikri124/retail-admin-app/internal/handler"
	"github.com/zikri124/retail-admin-app/internal/infrastructure"
	"github.com/zikri124/retail-admin-app/internal/repository"
	"github.com/zikri124/retail-admin-app/internal/router"
	"github.com/zikri124/retail-admin-app/internal/service"
)

func main() {
	g := gin.Default()

	gorm := infrastructure.NewGormPostgres()

	//prepare the order routes group
	orderRouteGroup := g.Group("/v1/orders")
	orderRepo := repository.NewOrderRepository(gorm)
	orderServices := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderServices)
	orderRouter := router.NewOrderRouter(orderRouteGroup, orderHandler)

	//mount routes group
	orderRouter.Mount()

	g.Run("127.0.0.1:3000")
}
