package services

import (
	"Taskmanager/EcommerceManager/common"
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Update product
// @Description Create product
// @Tags /api/product
// @Accept  json
// @Param product body models.Product true "Product data"
// @Produce  json
// @Router /api/product [post]
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

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete any product using their id
// @Tags /api/product
// @Produce  json
// @Router /api/product [delete]
func DeleteProduct(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print("[ERROR] converting product id failed", err)
		return
	}
	result := database.DB.Delete(&models.Product{}, productId)
	common.LogStatus(ctx, uint(productId), "", "deleted", result.Error, "product")
}

func ListOneProduct(ctx *gin.Context) {
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
		log.Printf("Error fetching product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	// Respond with the categories in JSON format
	ctx.JSON(http.StatusOK, products)
}

// ListProduct godoc
// @Summary List all products
// @Description Get a list of all products, including their variants.
// @Tags /api/product
// @Produce  json
// @Router /api/product [get]
func ListProduct(ctx *gin.Context) {

	var products []models.Product

	// Fetch categories with their child categories and products under each child category
	if err := database.DB.
		Preload("Variants"). // Preload variants under each child product
		Find(&products).Error; err != nil {
		log.Printf("Error fetching products: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	// Respond with the categories in JSON format
	ctx.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update any product using their id
// @Tags /api/product
// @Accept  json
// @Produce  json
// @Router /api/product [put]
func UpdateProduct(ctx *gin.Context) {
	productId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting product id", conv_err)
		return
	}

	var updates map[string]interface{}

	// Bind the JSON body to the updates map
	err := ctx.ShouldBindJSON(&updates)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Perform the update with a WHERE condition
	result := database.DB.Model(&models.Product{}).
		Where("id = ?", productId).
		Updates(updates)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})

}
