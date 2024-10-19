package handler

import (
	"Taskmanager/EcommerceManager/managers"

	"github.com/gin-gonic/gin"
)

func RegisterVariantApis(router *gin.Engine) {
	variantsGroup := router.Group("/api/variants")
	variantsGroup.POST("", managers.CreateVariant)
	variantsGroup.DELETE(":id", managers.DeleteVariant)
	variantsGroup.GET("", managers.ListVariant)
	variantsGroup.GET(":id", managers.ListOneVariant)
	variantsGroup.PUT(":id", managers.UpdateVariant)
}
