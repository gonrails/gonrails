package admin

import (
	"kalista/controllers/admin/revenue_reports"
	"kalista/controllers/admin/topic_reports"

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
		adminGroup.GET("/ping", ping)
		adminGroup.GET("/revenue_reports", revenue_reports.Index)
		adminGroup.GET("/topics_reports", topic_reports.Index)
	}
}
