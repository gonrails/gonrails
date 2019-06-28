package tenants

import (
	"github.com/gin-gonic/gin"
	"github.com/one-hole/gonrails/models"
	"github.com/one-hole/gonrails/serializers"
	. "github.com/one-hole/gonrails/serializers/tenant"
)

func Show(ctx *gin.Context) {
	var tenant = &models.Tenant{}
	models.DB().Model(tenant).First(tenant)

	ctx.JSON(200, gin.H{
		"data": serializers.SingleSerializer(&AdminTenantShowSerializer{}, tenant),
		"meta": 322,
	})
}
