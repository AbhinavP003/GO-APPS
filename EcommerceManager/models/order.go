package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderId    uint
	Status     bool
	Items      []uint `gorm:"type:json;default:null"`
	OrderTotal uint
}
