package tenants

import (
	"kalista/models"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	var tenant = &models.Tenant{}
	models.DB().Model(tenant).First(tenant)
	data, _ := models.Tenants(0, 10)
	ctx.JSON(200, gin.H{
		"data": data,
	})
}
