package tenants

import (
	"kalista/models"
	"kalista/serializers"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	var tenant = &models.Tenant{}
	models.DB().Model(tenant).First(tenant)

	call(ctx)

	data, _ := models.Tenants(0, 10)
	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func call(ctx *gin.Context) {
	var tenant = &models.Tenant{}
	models.DB().Model(tenant).First(tenant)

	serializers.SingleSerializer(tenant, &models.AdminTenantIndexSerializer{})
}

//{
//	data: {
//
//	},
//	meta: {
//
//	}
//}

//{
//	data: [
//	],
//	meta: {
//
//	}
//}
