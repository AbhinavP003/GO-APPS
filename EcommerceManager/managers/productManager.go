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
	productId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting product id", conv_err)
		return
	}
	var products []models.Product
	// Fetch categories with their child categories and products under each child category
	if err := database.DB.
		Preload("Variants").
		Where("id = ?", productId). // Preload variants under each child product
		First(&products).Error; err != nil {
		log.Printf("Error fetching products: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	// Respond with the categories in JSON format
	ctx.JSON(http.StatusOK, products)
}
func ListOneProduct(ctx *gin.Context) {

	var products []models.Product

	// Fetch categories with their child categories and products under each child category
	if err := database.DB.
		Preload("Variants"). // Preload variants under each child product
		First(&products).Error; err != nil {
		log.Printf("Error fetching product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	// Respond with the categories in JSON format
	ctx.JSON(http.StatusOK, products)
}
