package domain

type Item struct {
	Id          uint32 `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     uint32 `json:"order_id"`
}
