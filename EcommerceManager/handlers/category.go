package handler

import (
	services "Taskmanager/EcommerceManager/services"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryApis(router *gin.Engine) {
	categoryGroup := router.Group("/api/category")
	categoryGroup.POST("", services.CreateCategory)
	categoryGroup.DELETE(":id", services.DeleteCategory)
	categoryGroup.GET(":id", services.ListOneCategory)
	categoryGroup.GET("", services.ListCategory)
	categoryGroup.PUT(":id", services.UpdateCategory)
}
