package models

import (
	//	"database/sql/driver"
	//	"encoding/json"
	"gorm.io/gorm"
)

type Hierarchy struct {
	gorm.Model
	CategoryName    string `json:"category_name" gorm:"size:255"`
	ChildCategories string `json:"child_categories" gorm:"size:255;default:null"`
	//	ProductIDs      []uint 		  `json:prod"type:json;default:null"`
}

//type ProductIDList

//// Value marshals the ProductIDList to JSON for storing in the DB
//func (p ProductIDList) Value() (driver.Value, error) {
//	return json.Marshal(p) // Convert slice of uint to JSON
//}
//
//// Scan unmarshals JSON from the DB into a ProductIDList
//func (p *ProductIDList) Scan(value interface{}) error {
//	return json.Unmarshal(value.([]byte), &p) // Convert JSON from DB to []uint
//}
