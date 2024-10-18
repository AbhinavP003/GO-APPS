package handler

import (
	"Taskmanager/EcommerceManager/managers"

	"github.com/gin-gonic/gin"
)


func RegisterVariantApis(router *gin.Engine) {
	variantsGroup := router.Group("/api/product")
	variantsGroup.POST("", managers.CreateProduct)
	variantsGroup.DELETE(":id", managers.DeleteProduct)
	variantsGroup.GET("", managers.ListProduct)
}