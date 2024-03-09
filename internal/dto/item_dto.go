package dto

import "github.com/zikri124/retail-admin-app/internal/domain"

type ItemDto struct {
	LineItemId  uint32 `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (i *ItemDto) Transform(item domain.Item) {
	i.LineItemId = item.Id
	i.ItemCode = item.ItemCode
	i.Description = item.Description
	i.Quantity = item.Quantity
}
