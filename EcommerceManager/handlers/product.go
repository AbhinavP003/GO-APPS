package handler

import (
	"Taskmanager/EcommerceManager/services"

	"github.com/gin-gonic/gin"
)

func RegisterProductApis(router *gin.Engine) {
	prodcutGroup := router.Group("/api/product")
	prodcutGroup.POST("", services.CreateProduct)
	prodcutGroup.DELETE(":id", services.DeleteProduct)
	prodcutGroup.GET("", services.ListProduct)
	prodcutGroup.GET(":id", services.ListOneProduct)
	prodcutGroup.PUT(":id", services.UpdateProduct)
}
