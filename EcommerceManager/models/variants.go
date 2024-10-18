package models

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	Name          string `json:"name" gorm:"size:255"`
	MRP           uint   `json:"mrp"`
	DsicountPrice uint   `json:"discount_price"`
	Size          string `json:"size" gorm:"size:255"`
	Colour        string `json:"colour" gorm:"size:255"`
	Quantity      uint   `json:"quantity"`
	ProductId     uint   `json:"product_id"`
}
