package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"encoding/json"
	"io"
	"log"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	bodyAsByteArray, _ := io.ReadAll(ctx.Request.Body)
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bodyAsByteArray, &jsonMap)
	orderPlaced := placeOrder(jsonMap)
	if orderPlaced {
		ctx.JSON(http.StatusOK, gin.H{"msg": "Created Order"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Failed creating Order"})
}

func ListOrder(ctx *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func placeOrder(orderData map[string]interface{}) (orderPlaced bool) {
	variantId := orderData["variant_id"]
	quantity, ok := orderData["quantity"].(uint)
	if !ok {
		log.Print("Product input quantity not an integer ")
		return false
	}

	variant := models.Variant{}
	database.DB.First(&variant, variantId)
	if quantity > variant.Quantity {
		log.Printf("Selected stock quantity not available, remaining stock count: %d", variant.Quantity)
		return false
	}
	return true

}
