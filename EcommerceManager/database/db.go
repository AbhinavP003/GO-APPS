package database

import (
  "Taskmanager/EcommerceManager/models"
  "fmt"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

var DB *gorm.DB

func Initialise() {
  var err error
  fmt.Println("Initialsing db")
  DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
  fmt.Println("Initialised db")
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  DB.AutoMigrate(&models.Hierarchy{}, &models.Order{}, &models.Product{}, &models.Variant{})
}
