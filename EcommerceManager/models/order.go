package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status     string
	VariantId  string
	Quantity   uint
	OrderTotal uint
}
