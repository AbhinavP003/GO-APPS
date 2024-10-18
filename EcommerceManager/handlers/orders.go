package handler

import (
	"Taskmanager/EcommerceManager/managers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderApis(router *gin.Engine) {
	ordersGroup := router.Group("/api/order")
	ordersGroup.POST("", managers.CreateOrder)
	ordersGroup.DELETE(":id", managers.DeleteOrder)
	ordersGroup.GET("", managers.ListOrder)
}
