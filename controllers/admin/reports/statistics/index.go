/*
Author: 曾涛
Desc:   GET /admin/reports/statistics
Date:   2019-05-06
Email:  zengtao@risewinter.com
*/

package statistics

import (
	"kalista/controllers"
	"kalista/models"
	"kalista/serializers"
	"kalista/serializers/report"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	records, record := []*models.StatisticReports{}, &models.StatisticReports{}
	query := ctx.Request.URL.Query()
	record.FilterParams(controllers.QueryToMap(query)).Find(&records)

	resp, err := serializers.CollectionSerializer(&report.AdminStatisticReportIndexSerializer{}, records)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
