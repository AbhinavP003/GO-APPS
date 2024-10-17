package main

import (
	userHandler "Taskmanager/UserManager/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userHandler.RegisterUserApis(router)
	port := "3000"
	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
