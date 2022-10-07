package order

import (
	"assignment2/item"
	"time"
)

type Order struct {
	OrderID      int         `json:"order_id" gorm:"type:integer;primaryKey;auto_increment"`
	CustomerName string      `json:"customer_name" gorm:"type:varchar(50)"`
	Items        []item.Item `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderedAt    time.Time
}

type OrderRequest struct {
	OrderedAt    time.Time          `json:"orderedAt" binding:"required"`
	CustomerName string             `json:"customerName" binding:"required"`
	Items        []item.ItemRequest `json:"items" binding:"required"`
}

type OrderResponse struct {
	OrderID      int                 `json:"orderId"`
	CustomerName string              `json:"customerName"`
	OrderedAt    time.Time           `json:"orderedAt"`
	Items        []item.ItemResponse `json:"items" `
}
