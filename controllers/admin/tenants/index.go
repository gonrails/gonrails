package tenants

import (
	"github.com/one-hole/gonrails/models"
	"github.com/one-hole/gonrails/serializers"
	. "github.com/one-hole/gonrails/serializers/tenant"

	"github.com/gin-gonic/gin"
)

// Index GET - /admin/tenants
func Index(ctx *gin.Context) {
	var tenant = &models.Tenant{}
	models.DB().Model(tenant).First(tenant)

	data, _ := models.Tenants(0, 10)

	resp, err := serializers.CollectionSerializer(&AdminTenantIndexSerializer{}, data)

	if err != nil {
		ctx.JSON(400, gin.H{
			"data": err,
		})
	}

	ctx.JSON(200, gin.H{
		"data": resp,
	})
}

//[]*models.Tenant
