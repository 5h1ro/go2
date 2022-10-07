package item

type Item struct {
	ItemID      int    `json:"item_id" gorm:"type:integer;primaryKey;auto_increment"`
	ItemCode    string `json:"item_code" gorm:"type:varchar(10);"`
	Description string `json:"description" gorm:"type:text"`
	Quantity    int    `json:"quantity" gorm:"type:integer"`
	OrderID     int    `json:"order_id" gorm:"type:integer;not null"`
}

type ItemRequest struct {
	ItemID      int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" gorm:"type:integer"`
	OrderID     int    `json:"order_id"  binding:"required"`
}

type ItemResponse struct {
	ItemID      int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
