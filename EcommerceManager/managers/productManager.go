package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	productData := models.Product{}
	err := ctx.BindJSON(&productData)
	if err != nil {
		fmt.Println("error in binding post data", err)
		return
	}
	{
		database.DB.Create(&productData)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Created Product"})
	}
}

func DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	{
		database.DB.Delete(&models.Product{}, productId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Product deleted with id"})
	}
}

func ListProduct(ctx *gin.Context) {
	var products []models.Hierarchy
	database.DB.Find(&products)
	ctx.JSON(http.StatusOK, gin.H{"data": products})
}
