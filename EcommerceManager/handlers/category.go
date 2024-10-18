package handler

import (
	"Taskmanager/EcommerceManager/managers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryApis(router *gin.Engine) {
	categoryGroup := router.Group("/api/category")
	categoryGroup.POST("", managers.CreateCategory)
	categoryGroup.DELETE(":id", managers.DeleteCategory)
	categoryGroup.GET("", managers.ListCategory)
}
