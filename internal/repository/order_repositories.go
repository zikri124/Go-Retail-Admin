package repository

import (
	"context"

	"github.com/zikri124/retail-admin-app/internal/domain"
	"github.com/zikri124/retail-admin-app/internal/infrastructure"
	"gorm.io/gorm"
)

type OrderQuery interface {
	GetOrders(ctx context.Context) ([]domain.Order, error)
	CreateOrder(ctx context.Context, order *domain.Order) error
	UpdateOrder(ctx context.Context, order *domain.Order) error
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

func (o *orderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	db := o.db.GetConnection()

	if err := db.WithContext(ctx).Table("orders").Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) UpdateOrder(ctx context.Context, order *domain.Order) error {
	db := o.db.GetConnection()

	err := db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error

	if err != nil {
		return err
	}

	return nil
}
