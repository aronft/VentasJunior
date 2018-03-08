package models

import (
	"github.com/jinzhu/gorm"
)

// Product son los productos de la tienda
type Product struct {
	gorm.Model
	CategoryID  uint          `json:"categoryId" gorm:"not null"`
	Name        string        `json:"name" gorm:"not null; unique"`
	Price       float64       `json:"price" gorm:"not null;type:numeric(19,7)"`
	Description string        `json:"description"`
	OrderDetail []OrderDetail `json:"orderDatail,omitempty"`
}
