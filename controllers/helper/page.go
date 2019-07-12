package helper

import "github.com/gin-gonic/gin"

/*CurrentPage returns the current paginate page number, returns 1 as default
 */
func CurrentPage(ctx *gin.Context) int {
	return currentPage(ctx)
}

func currentPage(ctx *gin.Context) int {
	return 1
}
