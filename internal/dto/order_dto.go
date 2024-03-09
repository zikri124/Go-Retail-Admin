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

func (o *OrderDto) TransformToDto(order domain.Order) {
	o.Id = order.Id
	o.CustomerName = order.CustomerName
	o.OrderedAt = order.OrderedAt

	items := []ItemDto{}
	for _, v := range order.Items {
		item := ItemDto{}
		item.Transform(v)
		items = append(items, item)
	}

	o.Items = items
}

func (o *OrderDto) TransformToDomain(order domain.Order) domain.Order {
	order.CustomerName = o.CustomerName
	order.OrderedAt = o.OrderedAt

	return order
}
