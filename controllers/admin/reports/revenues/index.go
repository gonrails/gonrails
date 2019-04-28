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
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	var revenues = []*models.RevenueReport{}
	var report = &models.RevenueReport{}

	querys := ctx.Request.URL.Query()
	report.FilterParams(controllers.QueryToMap(querys)).Find(&revenues)

	//resp, err := serializers.CollectionSerializer()

	ctx.JSON(http.StatusOK, gin.H{
		"data": revenues,
	})
}

// Tenant
// Game
// Date
// Start
