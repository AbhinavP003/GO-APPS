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

func CreateVariant(ctx *gin.Context) {
	variantData := models.Variant{}
	err := ctx.BindJSON(&variantData)
	if err != nil {
		log.Print("[ERROR] error in binding variant data", err)
		return
	}
	result := database.DB.Create(&variantData)
	common.LogStatus(ctx, variantData.ID, variantData.Name, "created", result.Error, "variant")
}

func DeleteVariant(ctx *gin.Context) {
	variantId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print("[ERROR] error in converting variant id", err)
		return
	}
	result := database.DB.Delete(&models.Variant{}, variantId)
	common.LogStatus(ctx, uint(variantId), "", "deleted", result.Error, "variant")
}

func ListVariant(ctx *gin.Context) {
	var variants []models.Variant
	database.DB.Find(&variants)
	ctx.JSON(http.StatusOK, gin.H{"data": variants})
}
func ListOneVariant(ctx *gin.Context) {
	variantId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting variant id", conv_err)
		return
	}
	var variants []models.Variant
	database.DB.First(&variants, variantId)
	ctx.JSON(http.StatusOK, variants)
}

func UpdateVariant(ctx *gin.Context) {
	variantId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting variant id", conv_err)
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
	result := database.DB.Model(&models.Variant{}).
		Where("id = ?", variantId).
		Updates(updates)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update variant"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Variant updated successfully"})

}
