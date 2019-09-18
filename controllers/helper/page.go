package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	defaultSize = 10
)

/*CurrentPage returns the current paginate page number, returns 1 as default
 */
func CurrentPage(ctx *gin.Context) int {
	return currentPage(ctx)
}

// where create query. use GET method and put page into query params
func currentPage(ctx *gin.Context) int {
	if p, ok := ctx.GetQuery("page"); ok {
		page, _ := strconv.Atoi(p)
		return page
	}
	return 1 // Set Default
}

func CurrentSize(ctx *gin.Context) int {
	return currentSize(ctx)
}

func currentSize(ctx *gin.Context) int {
	if value, ok := ctx.GetQuery("size"); ok {
		size, _ := strconv.Atoi(value)
		return size
	}
	return defaultSize
}

// Page - paginate method for Controller use
func Page(db *gorm.DB, pageNum int) *gorm.DB {
	return page(db, pageNum)
}

func page(db *gorm.DB, pageNum int) *gorm.DB {
	return db.Offset((pageNum - 1) * defaultSize).Limit(defaultSize)
}

func PageWithSize(db *gorm.DB, pageNum, pageSize int) *gorm.DB {
	return pageWithSize(db, pageNum, pageSize)
}

func pageWithSize(db *gorm.DB, pageNum, pageSize int) *gorm.DB {
	return db.Offset( (pageNum - 1) * pageSize).Limit(pageSize)
}
