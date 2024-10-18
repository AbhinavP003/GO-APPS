package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status     bool   `json:"status"`
	Items      []uint `json:"items" gorm:"type:json;default:null"`
	OrderTotal uint   `json:"order_total" gorm:"size:255"`
}
