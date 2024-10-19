package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status     string   `json:"status"`
	Items      map[string]interface{} `json:"items" gorm:"type:json;default:null"`
	OrderTotal uint   `json:"order_total" gorm:"size:255"`
}
