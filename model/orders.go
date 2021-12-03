package model

type Orders struct {
	ID uint   `json:"id"`
	CustomerName string `json:"customerName"`
	OrderedAt string `json:"orderedAt"`
	Items      []Items `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}