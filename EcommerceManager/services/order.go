package services

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"encoding/json"
	"io"
	"log"
	"reflect"
	"strconv"

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
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully created order"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Failed creating Order"})
}

func ListOrder(ctx *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}
func ListOneOrder(ctx *gin.Context) {
	orderId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting order id", conv_err)
		return
	}
	var order []models.Order
	database.DB.First(&order, orderId)
	ctx.JSON(http.StatusOK, order)
}

func placeOrder(orderData map[string]interface{}) (orderPlaced bool) {
	// fetching required variant row
	variantId := orderData["variant_id"]
	variant := models.Variant{}
	database.DB.First(&variant, variantId)


	quantity, ok := orderData["quantity"].(float64)
	if !ok {
		log.Printf("Expected an integer for 'quantity', but got %v", reflect.TypeOf(orderData["quantity"]))
		return
	}
	if quantity > float64(variant.Quantity) {
		log.Printf("Selected stock quantity not available, remaining stock count: %d", variant.Quantity)
		return
	}
	// creating order in Orders table
	order := models.Order{}
	order.Status = "Accepted"
	totalPrice := variant.MRP * uint(quantity)
	order.OrderTotal = totalPrice
	order.Quantity = uint(quantity)
	result := database.DB.Create(&order)
	if result.Error != nil {
		log.Printf("Failed creating order %v", result.Error)
		return
	}

	// updating variant table with new stock quantity
	updatedQuantity := float64(variant.Quantity) - quantity
	variant_res := database.DB.Model(&variant).Where("id = ?", variantId).Update("quantity", updatedQuantity)
	if variant_res.Error != nil {
		log.Printf("Failed to update stock quantity: %v", variant_res.Error)
		return
	}
	if variant_res.RowsAffected == 0 {
		log.Printf("No variant found with id %d", variantId)
		return
	}

	return true

}
