package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductId     uint
	Name          string
	Description   string
	ImageUrl      string
	ChildVariants []uint `gorm:"type:json;default:null"`
}
