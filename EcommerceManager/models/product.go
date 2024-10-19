package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description string    `json:"description" gorm:"size:255"`
	CategoryID  uint      // Foreign key to Category
	Category    Category  `json:"-" gorm:"foreignKey:CategoryID"`
	Variants    []Variant `json:"variants,omitempty"`
}
