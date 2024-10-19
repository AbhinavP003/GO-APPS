package models

import "gorm.io/gorm"

// 	gorm.Model
// 	Name          string    `json:"name" gorm:"size:255"`
// 	Description   string    `json:"description" gorm:"size:255"`
// 	ImageUrl      string    `json:"image_url" gorm:"size:255"`
// 	ChildVariants []Variant `gorm:"foreignKey:ProductId"`
// }

// type Product struct {
// 	gorm.Model
// 	Name        string    `gorm:"not null"`
// 	Description string    `json:"description" gorm:"size:255"`
// 	CategoryID  uint      `gorm:"index"` // Foreign key to Category
// 	Category    Category  `gorm:"foreignKey:CategoryID"`
// 	Variants    []Variant `gorm:"foreignKey:ProductId"`
// }

type Product struct {
	gorm.Model
	Name        string   `gorm:"not null"`
	Description string   `json:"description" gorm:"size:255"`
	CategoryID  uint     // Foreign key to Category
	Category    Category `gorm:"foreignKey:CategoryID"`
	Variants    string
}
