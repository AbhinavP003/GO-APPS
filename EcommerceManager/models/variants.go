package models

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	VariantId     uint
	Name          string
	MRP           uint
	DsicountPrice uint
	Size          string
	Colour        string
	Quantity      string
}
