package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string    `json:"name" gorm:"size:255"`
	Description   string    `json:"description" gorm:"size:255"`
	ImageUrl      string    `json:"image_url" gorm:"size:255"`
	ChildVariants []Variant `gorm:"foreignKey:ProductId"`
}
