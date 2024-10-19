package models

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	Name          string
	MRP           uint    `gorm:"not null"`
	DiscountPrice string  `json:"discount_price"`
	Size          string  `json:"size" gorm:"size:255"`
	Colour        string  `json:"colour" gorm:"size:255"`
	Quantity      uint    `json:"quantity"`
	ProductId     uint    // Foreign key to Product
	Product       Product `json:"-" gorm:"foreignKey:ProductId"`
}
