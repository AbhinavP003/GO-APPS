package managers

import (
	"Taskmanager/EcommerceManager/common"
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
		order := models.Order{}
		order.Status = "Accepted"
		items := map[string]interface{}{"variantId": jsonMap["variant_id"], "quantity": jsonMap["quantity"], "price": jsonMap["mrp"]}
		order.Items = items
		result := database.DB.Create(&order)
		common.LogStatus(ctx, order.ID, "", "created", result.Error, "order")
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
		return
	}

	variant := models.Variant{}
	database.DB.First(&variant, variantId)
	if quantity > variant.Quantity {
		log.Printf("Selected stock quantity not available, remaining stock count: %d", variant.Quantity)
		return
	}
	updatedQuantity := variant.Quantity - quantity
	result := database.DB.Model(&variant).Where("id = ?", variantId).Update("quantity", updatedQuantity)
	if result.Error != nil {
		log.Printf("Failed to update stock quantity: %v", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		log.Printf("No variant found with id %d", variantId)
		return
	}
	return true

}
