package userHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserApis(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Grouped user api",
			})
		})
		userGroup.POST("/adduser", func(ctx *gin.Context){
			
		})
	}

}
