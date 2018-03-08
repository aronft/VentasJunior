package models

import (
	"github.com/jinzhu/gorm"
)

//Category es la categoria de productos
type Category struct {
	gorm.Model
	Name    string    `json:"name" gorm:"not null"`
	Product []Product `json:"product,omitempty"`
}
