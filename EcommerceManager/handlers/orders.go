package handler

import (
	"Taskmanager/EcommerceManager/services"

	"github.com/gin-gonic/gin"
)

func RegisterOrderApis(router *gin.Engine) {
	ordersGroup := router.Group("/api/order")
	ordersGroup.POST("", services.CreateOrder)
	ordersGroup.GET("", services.ListOrder)
	ordersGroup.GET(":id", services.ListOneOrder)
}
