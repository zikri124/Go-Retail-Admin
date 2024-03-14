package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zikri124/retail-admin-app/internal/handler"
	"github.com/zikri124/retail-admin-app/internal/infrastructure"
	"github.com/zikri124/retail-admin-app/internal/repository"
	"github.com/zikri124/retail-admin-app/internal/router"
	"github.com/zikri124/retail-admin-app/internal/service"

	_ "github.com/zikri124/retail-admin-app/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			RETAIL ADMIN API DOCUMENTATION
// @version		2.0
// @description	api doc for golang bootcamp hackativ8 x kominfo
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
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

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run("127.0.0.1:3000")
}
