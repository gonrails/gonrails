package admin

import (
	"kalista/controllers/admin/reports"
	"kalista/controllers/admin/tenants"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(router *gin.RouterGroup) {
	adminGroup := router.Group("/admin")
	{
		reports.ApplyRoutes(adminGroup)
		adminGroup.POST("/tenants", tenants.Create)
	}
}

func currentEmployee() {

}

func authenticate(ctx *gin.Context) bool {
	return true
}
