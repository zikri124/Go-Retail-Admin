package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Id          uint32     `json:"id"`
	ItemCode    string     `json:"item_code"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	OrderId     uint32     `json:"order_id"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (i *Item) BeforeCreate(db *gorm.DB) (err error) {
	if i.Id == 0 {
		i.Id = uuid.New().ID()
	}
	return
}
