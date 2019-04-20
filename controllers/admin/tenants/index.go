package tenants

import (
	"fmt"
	"kalista/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	var tenant = &models.Tenant{}

	models.DB().Model(tenant).First(tenant)

	fmt.Println(tenant.Serialize())

	data, _ := models.Tenants(0, 10)
	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func currentPage(ctx *gin.Context) int {
	page, ok := ctx.GetQuery("page")

	if !ok {
		page = "1" // Default Page
	}

	currentPage, _ := strconv.Atoi(page)
	return currentPage
}
