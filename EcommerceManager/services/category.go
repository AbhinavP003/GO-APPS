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

// CreateCategory godoc
// @Summary Create category
// @Description Create category 
// @Tags /api/category
// @Param product body models.Category true "Category data"
// @Accept  json
// @Produce  json
// @Router /api/category [post]
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

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete any category using their id
// @Tags /api/category
// @Produce  json
// @Router /api/category [delete]
func DeleteCategory(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print("[ERROR] error in converting category id", err)
		return
	}
	result := database.DB.Delete(&models.Category{}, categoryId)
	common.LogStatus(ctx, uint(categoryId), "", "deleted", result.Error, "category")
}

// ListCategory godoc
// @Summary List all categories
// @Description List categories
// @Tags /api/category
// @Produce  json
// @Router /api/category [get]
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

func ListOneCategory(ctx *gin.Context) {
	categoryId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting category id", conv_err)
		return
	}
	var category models.Category
	err := database.DB.
		Preload("ChildCategories.ChildCategories").        // Preload nested child categories
		Preload("ChildCategories.ChildProducts").          // Preload products under each child category
		Preload("ChildCategories.ChildProducts.Variants"). // Preload variants under each child product
		Where("id = ?", categoryId).                       // Fetch the category by ID
		First(&category).Error

	if err != nil {
		log.Printf("Error fetching category: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		return
	}

	// Return the fetched category with nested relations
	ctx.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update any category using their id
// @Tags /api/category
// @Accept  json
// @Produce  json
// @Router /api/category [put]
func UpdateCategory(ctx *gin.Context) {
	categoryId, conv_err := strconv.Atoi(ctx.Param("id"))
	if conv_err != nil {
		log.Print("[ERROR] error in converting category id", conv_err)
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
	result := database.DB.Model(&models.Category{}).
		Where("id = ?", categoryId).
		Updates(updates)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})

}
