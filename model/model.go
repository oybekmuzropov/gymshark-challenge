package model

type Pack struct {
	PackSize int `json:"pack_size"`
	Count    int `json:"count"`
}

type CalculatePackReq struct {
	OrderedItemsCount int `json:"ordered_items_count" validate:"required,gt=0"`
}
