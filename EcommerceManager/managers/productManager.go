package managers

import (
	"Taskmanager/EcommerceManager/common"
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	productData := models.Product{}
	err := ctx.BindJSON(&productData)
	if err != nil {
		log.Print("[ERROR] error in binding product data", err)
		return
	}
	result := database.DB.Create(&productData)
	common.LogStatus(ctx, productData.ID, productData.Name, "created", result.Error, "product")
}

func DeleteProduct(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print("[ERROR] converting product id failed", err)
		return
	}
	result := database.DB.Delete(&models.Product{}, productId)
	common.LogStatus(ctx, uint(productId), "", "deleted", result.Error, "product")
}

func ListProduct(ctx *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	ctx.JSON(http.StatusOK, gin.H{"data": products})
}
