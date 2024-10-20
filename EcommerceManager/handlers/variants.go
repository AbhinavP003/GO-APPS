package handler

import (
	"Taskmanager/EcommerceManager/services"

	"github.com/gin-gonic/gin"
)

func RegisterVariantApis(router *gin.Engine) {
	variantsGroup := router.Group("/api/variants")
	variantsGroup.POST("", services.CreateVariant)
	variantsGroup.DELETE(":id", services.DeleteVariant)
	variantsGroup.GET("", services.ListVariant)
	variantsGroup.GET(":id", services.ListOneVariant)
	variantsGroup.PUT(":id", services.UpdateVariant)
}
