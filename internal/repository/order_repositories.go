package repository

import (
	"context"

	"github.com/zikri124/retail-admin-app/internal/domain"
	"github.com/zikri124/retail-admin-app/internal/infrastructure"
)

type OrderQuery interface {
	GetOrders(ctx context.Context) ([]domain.Order, error)
}

type orderRepository struct {
	db infrastructure.GormPostgres
}

func NewOrderRepository(db infrastructure.GormPostgres) OrderQuery {
	return &orderRepository{db: db}
}

func (o *orderRepository) GetOrders(ctx context.Context) ([]domain.Order, error) {
	db := o.db.GetConnection()
	orders := []domain.Order{}

	err := db.
		WithContext(ctx).
		Table("orders").
		Preload("Items").
		Find(&orders).
		Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}