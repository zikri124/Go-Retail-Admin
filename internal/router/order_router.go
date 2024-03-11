package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zikri124/retail-admin-app/internal/handler"
)

type OrderRouter interface {
	Mount()
}

type orderRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.OrderHandler
}

func NewOrderRouter(v *gin.RouterGroup, handler handler.OrderHandler) OrderRouter {
	return &orderRouterImpl{v: v, handler: handler}
}

func (o *orderRouterImpl) Mount() {
	o.v.GET("", o.handler.GetUsers)
	o.v.POST("", o.handler.CreateOrder)
	o.v.PUT("/:id", o.handler.UpdateOrder)
}
