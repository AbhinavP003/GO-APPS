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
