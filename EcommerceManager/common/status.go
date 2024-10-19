package common

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogStatus(ctx *gin.Context, id uint, name string, operation string, err error, modelName string) {
	if err != nil {
		errMsg := fmt.Sprintf("[ERROR] Failed to %s %s with id: %d and name: %s: error - %v", operation, modelName, id, name, err)
		log.Print(errMsg)
		ctx.JSON(http.StatusBadGateway, gin.H{"msg": errMsg})
		return
	}
	msg := fmt.Sprintf("Successfully %s %s with id: %d and name: %s", operation, modelName, id, name)
	ctx.JSON(http.StatusOK, gin.H{"msg": msg})
}
