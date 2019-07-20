package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

// Page - paginate method for cnotroller use
func Page(db *gorm.DB, pageNum int) *gorm.DB {
	return page(db, pageNum)
}

func page(db *gorm.DB, pageNum int) *gorm.DB {
	return db.Offset((pageNum - 1) * 10)
}
