package dto

import (
	"time"

	"github.com/zikri124/retail-admin-app/internal/domain"
)

type OrderDto struct {
	Id           uint32    `json:"orderId"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []ItemDto `json:"items"`
}

func (o *OrderDto) TransformToDto(order *domain.Order) {
	if order.Id != 0 {
		o.Id = order.Id
	}
	o.CustomerName = order.CustomerName
	o.OrderedAt = order.OrderedAt

	items := []ItemDto{}
	for _, v := range order.Items {
		item := ItemDto{}
		item.TransformToDto(v)
		items = append(items, item)
	}

	o.Items = items
}

func (o *OrderDto) TransformToDomain() *domain.Order {
	order := domain.Order{}

	if o.Id != 0 {
		order.Id = o.Id
	} else {
		o.Id = order.Id
	}
	order.CustomerName = o.CustomerName
	order.OrderedAt = o.OrderedAt

	items := []domain.Item{}
	for _, v := range o.Items {
		item := domain.Item{}
		item = *v.TransformToDomain()
		items = append(items, item)
	}

	order.Items = items

	return &order
}
