package models

import (
	"gorm.io/gorm"
)

type Hierarchy struct {
	gorm.Model
	CategoryName    string `json:"category_name" gorm:"size:255"`
	ChildCategories string `json:"child_categories" gorm:"size:255;default:null"`
}

