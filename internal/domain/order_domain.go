package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id           uint32     `json:"id"`
	CustomerName string     `json:"customer_name"`
	OrderedAt    time.Time  `json:"ordered_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Items        []Item     `json:"items" gorm:"foreignKey:OrderId;references:Id"`
}

func (o *Order) BeforeCreate(db *gorm.DB) (err error) {
	if o.Id == 0 {
		o.Id = uuid.New().ID()
	}
	return
}
