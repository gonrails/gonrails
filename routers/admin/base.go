package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/one-hole/gonrails/controllers/admin/tenants"
	"github.com/one-hole/gonrails/routers/admin/reports"
)

func ApplyRoutes(router *gin.RouterGroup) {
	adminGroup := router.Group("/admin")
	{
		reports.ApplyRoutes(adminGroup)
		adminGroup.GET("/tenants", tenants.Index)
		adminGroup.GET("/tenants/:id", tenants.Show)
		adminGroup.POST("/tenants", tenants.Create)
	}
}
