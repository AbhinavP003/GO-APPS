package main

import (
	"Taskmanager/EcommerceManager/database"
	handler "Taskmanager/EcommerceManager/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialise()
}

func main() {
	router := gin.Default()
	handler.RegisterHierarchyApis(router)
	handler.RegisterProductApis(router)
	handler.RegisterVariantApis(router)
	handler.RegisterOrderApis(router)
	err := router.Run()
	if err != nil {
		return
	}
}
