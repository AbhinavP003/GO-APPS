package main

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialise()
}

func main() {
	router := gin.Default()
	hierarchyHandler.RegisterHierarchyApis(router)
	err := router.Run()
	if err != nil {
		return
	}
}
