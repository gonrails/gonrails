package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
}

func DoResponse() {

}

func CurrentPage(ctx *gin.Context) int {
	page, ok := ctx.GetQuery("page")

	if !ok {
		page = "1" // Default Page
	}

	currentPage, _ := strconv.Atoi(page)
	return currentPage
}

func currentEmployee() {

}

func authenticate(ctx *gin.Context) bool {
	return true
}

// 需要有一些通用的方法

// currentPage
// filter 方法
