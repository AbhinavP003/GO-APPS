package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName    string      `json:"category_name" gorm:"size:255"`
	ChildCategories []Category `json:"child_categories" gorm:"foreignKey:ID"`
	Products        []Product   `json:"products" gorm:"foreignKey:ID"`
}
