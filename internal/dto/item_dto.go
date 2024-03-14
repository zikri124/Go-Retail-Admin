package dto

import "github.com/zikri124/retail-admin-app/internal/domain"

type ItemDto struct {
	LineItemId  uint32 `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ItemReqDto struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (i *ItemDto) TransformToDto(item domain.Item) {
	if item.Id != 0 {
		i.LineItemId = item.Id
	}
	i.ItemCode = item.ItemCode
	i.Description = item.Description
	i.Quantity = item.Quantity
}

func (i *ItemDto) TransformToDomain() *domain.Item {
	item := domain.Item{}

	if i.LineItemId != 0 {
		item.Id = i.LineItemId
	} else {
		i.LineItemId = item.Id
	}

	item.ItemCode = i.ItemCode
	item.Description = i.Description
	item.Quantity = i.Quantity

	return &item
}
