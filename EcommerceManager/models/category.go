package models

import "gorm.io/gorm"

// type Category struct {
// 	gorm.Model
// 	CategoryName    string     `json:"category_name" gorm:"not null"`
// 	ParentID        *uint      `gorm:"index"` // Self-referencing foreign key
// 	Parent          *Category  `gorm:"foreignKey:ParentID"`
// 	ChildCategories []Subcategory `json:"child_categories" gorm:"foreignKey:ParentID"` // To load child categories
// 	Products        []Product  `gorm:"foreignKey:CategoryID"`                       // Relationship with products
// }

// type Subcategory Category

type Category struct {
	gorm.Model
	Name    string     `json:"category_name" gorm:"size:255"`
	ParentCategory  *Category `gorm:"foreignKey:ParentId"`
	ParentId        *uint
	ChildCategories string
	ChildProducts        string
}
