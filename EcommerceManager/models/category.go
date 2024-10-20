package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name            string    `json:"category_name" gorm:"size:255"`
	ParentCategory  *Category `json:"-" gorm:"foreignKey:ParentId"`
	ParentId        *uint
	ChildCategories []Category `json:"child_categories,omitempty" gorm:"foreignKey:ParentId"`
	ChildProducts   []Product  `json:"child_products,omitempty"`
	Variants        []Variant  `json:"variants,omitempty" gorm:"foreignKey:ProductId"`
}
