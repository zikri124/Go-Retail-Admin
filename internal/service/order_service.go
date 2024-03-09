package service

import (
	"context"

	"github.com/zikri124/retail-admin-app/internal/domain"
	"github.com/zikri124/retail-admin-app/internal/repository"
)

type OrderService interface {
	GetOrders(ctx context.Context) ([]domain.Order, error)
}

type orderServiceImpl struct {
	repo repository.OrderQuery
}

func NewOrderService(repo repository.OrderQuery) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (o *orderServiceImpl) GetOrders(ctx context.Context) ([]domain.Order, error) {
	orders, err := o.repo.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
