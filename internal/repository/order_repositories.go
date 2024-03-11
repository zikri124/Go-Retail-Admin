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
	IsOrderExist(ctx context.Context, orderId uint32) (bool, error)
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
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Select("*").Where("deleted_at is null")
		}).
		Where("deleted_at is null").
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

func (o *orderRepository) IsOrderExist(ctx context.Context, orderId uint32) (bool, error) {
	db := o.db.GetConnection()

	var isExist bool
	err := db.
		WithContext(ctx).
		Table("orders").
		Select("count(*) > 0").
		Where("deleted_at is null and id = ?", orderId).
		Find(&isExist).
		Error

	return isExist, err

}
