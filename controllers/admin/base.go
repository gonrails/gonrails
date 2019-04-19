package admin

import (
	"github.com/gin-gonic/gin"
	"kalista/controllers/admin/reports"
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
	}
}

func currentEmployee() {

}

func authenticate(ctx *gin.Context) bool {
	return true
}
