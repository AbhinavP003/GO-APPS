package handler

import (
	"Taskmanager/EcommerceManager/managers"

	"github.com/gin-gonic/gin"
)


func RegisterOrderApis(router *gin.Engine) {
	ordersGroup := router.Group("/api/product")
	ordersGroup.POST("", managers.CreateProduct)
	ordersGroup.DELETE(":id", managers.DeleteProduct)
	ordersGroup.GET("", managers.ListProduct)
}