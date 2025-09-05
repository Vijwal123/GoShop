package models

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey"`
	UserID    uint        `json:"user_id"`
	Total     float64     `json:"total"`
	CreatedAt time.Time
	Items     []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
