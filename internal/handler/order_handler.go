package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikri124/retail-admin-app/internal/dto"
	"github.com/zikri124/retail-admin-app/internal/service"
	"github.com/zikri124/retail-admin-app/pkg/response"
)

type OrderHandler interface {
	GetUsers(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
}

type orderHandlerImpl struct {
	svc service.OrderService
}

func NewOrderHandler(svc service.OrderService) OrderHandler {
	return &orderHandlerImpl{svc: svc}
}

func (o *orderHandlerImpl) GetUsers(ctx *gin.Context) {
	orders, err := o.svc.GetOrders(ctx)
	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.SetSuccessResponse(ctx, http.StatusOK, orders)
}

func (o *orderHandlerImpl) CreateOrder(ctx *gin.Context) {
	orderDto := dto.OrderDto{}

	err := ctx.ShouldBindJSON(&orderDto)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	err = o.svc.CreateOrder(ctx, &orderDto)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.SetSuccessResponse(ctx, http.StatusCreated, orderDto)
	return
}
