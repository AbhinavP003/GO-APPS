package hierarchyHandler

import (
    "Taskmanager/EcommerceManager/managers"
    "github.com/gin-gonic/gin"
)

func RegisterHierarchyApis(router *gin.Engine) {
    hierarchyGroup := router.Group("/api/hierarchy")
    hierarchyGroup.POST("", managers.CreateHierarchy)
    //    hierarchyGroup.DELETE(":id", managers.DeleteHierarchy)
    hierarchyGroup.GET("", managers.ListHierarchy)
}
