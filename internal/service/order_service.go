package service

import (
	"context"

	"github.com/zikri124/retail-admin-app/internal/dto"
	"github.com/zikri124/retail-admin-app/internal/repository"
)

type OrderService interface {
	GetOrders(ctx context.Context) ([]dto.OrderDto, error)
	CreateOrder(ctx context.Context, orderDto *dto.OrderDto) error
	UpdateOrder(ctx context.Context, orderDto *dto.OrderDto) (*dto.OrderDto, error)
}

type orderServiceImpl struct {
	repo repository.OrderQuery
}

func NewOrderService(repo repository.OrderQuery) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (o *orderServiceImpl) GetOrders(ctx context.Context) ([]dto.OrderDto, error) {
	orders, err := o.repo.GetOrders(ctx)

	ordersDto := []dto.OrderDto{}
	for _, order := range orders {
		orderDto := dto.OrderDto{}
		orderDto.TransformToDto(&order)
		ordersDto = append(ordersDto, orderDto)
	}

	if err != nil {
		return nil, err
	}

	return ordersDto, nil
}

func (o *orderServiceImpl) CreateOrder(ctx context.Context, orderDto *dto.OrderDto) error {
	order := orderDto.TransformToDomain()
	err := o.repo.CreateOrder(ctx, order)
	orderDto.Id = order.Id

	orderDto.TransformToDto(order)

	return err
}

func (o *orderServiceImpl) UpdateOrder(ctx context.Context, orderDto *dto.OrderDto) (*dto.OrderDto, error) {
	order := orderDto.TransformToDomain()

	err := o.repo.UpdateOrder(ctx, order)

	return orderDto, err
}
