package reports

import (
	"kalista/controllers/admin/reports/revenues"
	"kalista/controllers/admin/reports/statistics"
	"kalista/controllers/admin/reports/topics"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.RouterGroup) {
	reportGroup := router.Group("/reports")
	{
		reportGroup.GET("/revenues", revenues.Index)
		reportGroup.GET("/topics", topics.Index)
		reportGroup.GET("/statistics", statistics.Index)
	}
}
