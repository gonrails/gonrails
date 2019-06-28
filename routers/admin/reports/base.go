package reports

import (
	"github.com/one-hole/gonrails/controllers/admin/reports/revenues"
	"github.com/one-hole/gonrails/controllers/admin/reports/statistics"
	"github.com/one-hole/gonrails/controllers/admin/reports/topics"

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

type myInt struct {
	value  int
	inited bool
}
