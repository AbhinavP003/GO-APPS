package handler

import (
	"Taskmanager/EcommerceManager/managers"

	"github.com/gin-gonic/gin"
)


func RegisterProductApis(router *gin.Engine) {
	prodcutGroup := router.Group("/api/product")
	prodcutGroup.POST("", managers.CreateProduct)
	prodcutGroup.DELETE(":id", managers.DeleteProduct)
	prodcutGroup.GET("", managers.ListProduct)
}