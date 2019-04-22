package admin

import (
	"kalista/controllers/admin/tenants"
	"kalista/routers/admin/reports"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.RouterGroup) {
	adminGroup := router.Group("/admin")
	{
		reports.ApplyRoutes(adminGroup)
		adminGroup.POST("/tenants", tenants.Create)
		adminGroup.GET("/tenants", tenants.Index)
	}
}
