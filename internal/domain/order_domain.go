package domain

import (
	"time"
)

type Order struct {
	Id           uint32    `json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId;references:Id"`
}
