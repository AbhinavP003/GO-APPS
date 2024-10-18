package main

import (
	"Taskmanager/EcommerceManager/database"
	handler "Taskmanager/EcommerceManager/handlers"
	"Taskmanager/EcommerceManager/logutil"
	"log"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialise()
	logutil.Initialise()
}

func main() {
	router := gin.Default()
	handler.RegisterHierarchyApis(router)
	handler.RegisterProductApis(router)
	handler.RegisterVariantApis(router)
	handler.RegisterOrderApis(router)
	err := router.Run()
	if err != nil {
		log.Print("[ERROR] failed to start server: ", err)
		return
	}
}
