package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	orderData := models.Order{}
	err := ctx.BindJSON(&orderData)
	if err != nil {
		log.Print("[ERROR] error in binding post data", err)
		return
	}
	{
		database.DB.Create(&orderData)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Created Order"})
	}
}

func DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	{
		database.DB.Delete(&models.Order{}, orderId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Order deleted with id"})
	}
}

func ListOrder(ctx *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}
