package models

import (
	"github.com/jinzhu/gorm"
)

// OrderDetail es el detalle de la orden
type OrderDetail struct {
	gorm.Model
	ProductID uint    `json:"productId" gorm:"not null"`
	OrderID   uint    `json:"orderId" gorm:"not null"`
	Quantity  uint    `json:"quantity" gorm:"not null"`
	Subtotal  float64 `json:"subtotal" gorm:"not null;type:numeric(19,7)"`
}
