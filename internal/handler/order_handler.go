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

// ShowAllOrders godoc
//
//	@Summary		Show orders list
//	@Description	list all orders with their items
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.SuccessRes
//	@Failure		400		{object}	response.ErrorRes
//	@Failure		404		{object}	response.ErrorRes
//	@Failure		500		{object}	response.ErrorRes
//	@Router			/v1/orders [get]
func (o *orderHandlerImpl) GetUsers(ctx *gin.Context) {
	orders, err := o.svc.GetOrders(ctx)
	if err != nil {
		response.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(ctx, http.StatusOK, orders)
}

// CreateNewOrder godoc
//
//	@Summary		Create a new order data
//	@Description	will save the new order data with their items to db
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		dto.OrderCreateDto	true	"New Order"
//	@Success		201		{object}	response.SuccessRes
//	@Failure		400		{object}	response.ErrorRes
//	@Failure		404		{object}	response.ErrorRes
//	@Failure		500		{object}	response.ErrorRes
//	@Router			/v1/orders [post]
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

// EditAnOrderData godoc
//
//	@Summary		Edit an order data
//	@Description	will get the body, and change order data with id and all items inside it
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Order ID"
//	@Param			order	body		dto.OrderUpdateDto	true	"New condition order"
//	@Success		200		{object}	response.SuccessRes
//	@Failure		400		{object}	response.ErrorRes
//	@Failure		404		{object}	response.ErrorRes
//	@Failure		500		{object}	response.ErrorRes
//	@Router			/v1/orders/{id} [put]
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

	isExist, err := o.svc.IsOrderExist(ctx, uint32(orderId))

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

// DeleteAnOrderData godoc
//
//	@Summary		Delete an order data
//	@Description	will soft delete order data by id and all its items
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200		{object}	response.SuccessRes
//	@Failure		400		{object}	response.ErrorRes
//	@Failure		404		{object}	response.ErrorRes
//	@Failure		500		{object}	response.ErrorRes
//	@Router			/v1/orders/{id} [delete]
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
