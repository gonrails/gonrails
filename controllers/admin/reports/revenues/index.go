/*
Author: 曾涛
Desc:   GET /admin/revenues
Date:   2019-04-26
Email:  zengtao@risewinter.com
*/

package revenues

import (
	"kalista/controllers"
	"kalista/models"
	"kalista/serializers"
	reports "kalista/serializers/report"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index - GET /admin/reports/revenues
func Index(ctx *gin.Context) {
	var revenues = []*models.RevenueReport{}
	var report = &models.RevenueReport{}

	querys := ctx.Request.URL.Query()
	report.FilterParams(controllers.QueryToMap(querys)).Find(&revenues)

	resp, err := serializers.CollectionSerializer(&reports.AdminRevenueReportIndexSerializer{}, revenues)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
		"meta": indexMeta(),
	})
}

func indexMeta() map[string]interface{} {
	return nil
}

// Tenant
// Game
// Date
// Start
