package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHierarchy(ctx *gin.Context) {
	hierarchyData := models.Hierarchy{}
	err := ctx.BindJSON(&hierarchyData)
	if err != nil {
		log.Print("[ERROR] error in binding post data", err)
		return
	}
	{
		database.DB.Create(&hierarchyData)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Created hierarchy"})
	}
}

func DeleteHierarchy(ctx *gin.Context) {
	hierarchyId := ctx.Param("id")
	{
		database.DB.Delete(&models.Hierarchy{}, hierarchyId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Hierarchy deleted with id"})
	}
}

func ListHierarchy(ctx *gin.Context) {
	var hierarchies []models.Hierarchy
	database.DB.Find(&hierarchies)
	ctx.JSON(http.StatusOK, gin.H{"Hierarchies": hierarchies})
}
