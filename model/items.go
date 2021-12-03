package model

type Items struct {
	ID        uint   `json:"id"`
	ItemCode string `json:"itemCode"`
	Description  string `json:"description"`
	Quantity      uint `json:"quantity"`
	OrderId uint
}