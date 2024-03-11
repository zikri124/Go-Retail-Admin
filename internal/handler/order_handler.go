package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zikri124/retail-admin-app/internal/dto"
	"github.com/zikri124/retail-admin-app/internal/service"
	"github.com/zikri124/retail-admin-app/pkg/response"
)

type OrderHandler interface {
	GetUsers(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
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
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(ctx, http.StatusOK, orders)
}

func (o *orderHandlerImpl) CreateOrder(ctx *gin.Context) {
	orderDto := dto.OrderDto{}

	err := ctx.ShouldBindJSON(&orderDto)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = o.svc.CreateOrder(ctx, &orderDto)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(ctx, http.StatusCreated, orderDto)
}

func (o *orderHandlerImpl) UpdateOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if orderId == 0 || err != nil {
		response.SetErrorResponse(ctx, http.StatusBadRequest, "invalid required param")
		return
	}

	orderDto := dto.OrderDto{}

	orderDto.Id = uint32(orderId)

	err = ctx.ShouldBindJSON(&orderDto)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	isExist, err := o.svc.IsOrderExist(ctx, orderDto.Id)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !isExist {
		response.SetErrorResponse(ctx, http.StatusNotFound, "Order not found")
		return
	}

	order, err := o.svc.UpdateOrder(ctx, &orderDto)

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(ctx, http.StatusOK, order)
}

func (o *orderHandlerImpl) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if orderId == 0 || err != nil {
		response.SetErrorResponse(ctx, http.StatusBadRequest, "invalid required param")
		return
	}

	isExist, err := o.svc.IsOrderExist(ctx, uint32(orderId))

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !isExist {
		response.SetErrorResponse(ctx, http.StatusNotFound, "Order not found")
		return
	}

	err = o.svc.DeleteOrder(ctx, uint32(orderId))

	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(ctx, http.StatusOK, "Order deleted")
}
