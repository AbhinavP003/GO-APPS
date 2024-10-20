package main

import (
	"Taskmanager/EcommerceManager/database"
	handler "Taskmanager/EcommerceManager/handlers"
	"Taskmanager/EcommerceManager/logutil"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Taskmanager/EcommerceManager/docs"

	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialise()
	logutil.Initialise()
}

func main() {
	sqlDb, _ := database.DB.DB()
	defer sqlDb.Close()
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handler.RegisterCategoryApis(router)
	handler.RegisterProductApis(router)
	handler.RegisterVariantApis(router)
	handler.RegisterOrderApis(router)
	err := router.Run()
	if err != nil {
		log.Print("[ERROR] failed to start server: ", err)
		return
	}
}
