package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/* RequestString provide a method to generate request info.
   We can use it with type APIError's Request attribute
   eg: GET: /home
 */
func RequestString(c *gin.Context) string {
	return fmt.Sprintf("%s: %s", c.Request.Method, c.Request.RequestURI)
}