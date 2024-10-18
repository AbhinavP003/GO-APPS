package models

import (
	"gorm.io/gorm"
)

type Hierarchy struct {
	gorm.Model
	CategoryName    string      `json:"category_name" gorm:"size:255"`
	ChildCategories []Hierarchy `json:"child_categories" gorm:"foreignKey:ID"`
	Products        []Product   `json:"products" gorm:"foreignKey:ID"`
}
