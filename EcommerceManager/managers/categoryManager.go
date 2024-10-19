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

func CreateCategory(ctx *gin.Context) {
	categoryData := models.Category{}
	err := ctx.BindJSON(&categoryData)
	if err != nil {
		log.Print("[ERROR] error in binding category data", err)
		return
	}
	result := database.DB.Create(&categoryData)
	common.LogStatus(ctx, categoryData.ID, categoryData.Name, "created", result.Error, "category")
}

func DeleteCategory(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print("[ERROR] error in converting variant id", err)
		return
	}
	result := database.DB.Delete(&models.Category{}, categoryId)
	common.LogStatus(ctx, uint(categoryId), "", "deleted", result.Error, "category")
}

//	func ListCategory(ctx *gin.Context) {
//		var categories []models.Category
//		database.DB.Find(&categories)
//		ctx.JSON(http.StatusOK, gin.H{"Categories": categories})
//	}
func ListCategory(ctx *gin.Context) {
	var categories []models.Category

	// Fetch categories with their child categories and products under each child category
	if err := database.DB.
		Preload("ChildCategories.ChildCategories").        // Preload nested child categories
		Preload("ChildCategories.ChildProducts").          // Preload products under each child category
		Preload("ChildCategories.ChildProducts.Variants"). // Preload variants under each child product
		Where("parent_id IS NULL").                        // Fetch only top-level categories
		Find(&categories).Error; err != nil {
		log.Printf("Error fetching categories: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	// Respond with the categories in JSON format
	ctx.JSON(http.StatusOK, categories)
}
