package managers

import (
	"Taskmanager/EcommerceManager/database"
	"Taskmanager/EcommerceManager/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHierarchy(ctx *gin.Context) {
	hierarchyData := models.Hierarchy{}
	err := ctx.BindJSON(&hierarchyData)
	if err != nil {
		fmt.Println("error in binding post data", err)
		return
	}
	{
		fmt.Println(hierarchyData)
		database.DB.Create(&hierarchyData)
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}

func DeleteHierarchy(ctx *gin.Context) {
	hierarchyId := ctx.Param("id")
	{
		database.DB.Delete(&models.Hierarchy{}, hierarchyId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "user deleted"})
	}
}

func ListHierarchy(ctx *gin.Context) {
	hierarchies := database.DB.Find(&models.Hierarchy{})
	{
		ctx.JSON(http.StatusOK, hierarchies)
	}
}
