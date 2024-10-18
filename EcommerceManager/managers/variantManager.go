package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateVariant(ctx *gin.Context) {
	variantData := models.Variant{}
	err := ctx.BindJSON(&variantData)
	if err != nil {
		fmt.Println("error in binding post data", err)
		return
	}
	{
		database.DB.Create(&variantData)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Created Variant"})
	}
}

func DeleteVariant(ctx *gin.Context) {
	variantId := ctx.Param("id")
	{
		database.DB.Delete(&models.Variant{}, variantId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Variant deleted with id"})
	}
}

func ListVariant(ctx *gin.Context) {
	var variants []models.Variant
	database.DB.Find(&variants)
	ctx.JSON(http.StatusOK, gin.H{"data": variants})
}
