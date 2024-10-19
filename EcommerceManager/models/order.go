package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status     string
	VariantId  string
	Price      uint
	Quantity   uint
	OrderTotal uint
}
