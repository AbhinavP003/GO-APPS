package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	categoryData := models.Category{}
	err := ctx.BindJSON(&categoryData)
	if err != nil {
		log.Print("[ERROR] error in binding post data", err)
		return
	}
	{
		database.DB.Create(&categoryData)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Created category"})
	}
}

func DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	{
		database.DB.Delete(&models.Category{}, categoryId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Category deleted"})
	}
}

func ListCategory(ctx *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)
	ctx.JSON(http.StatusOK, gin.H{"Categories": categories})
}
