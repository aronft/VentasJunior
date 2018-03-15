package models

import (
	"github.com/jinzhu/gorm"
)

// Order es la orden de compra o venta
type Order struct {
	gorm.Model
	UserID      uint          `json:"userId"`
	Total       float64       `json:"total" gorm:"not null;type:numeric(19,7)"`
	OrderDetail []OrderDetail `json:"orderDetail, omitempty"`
}
