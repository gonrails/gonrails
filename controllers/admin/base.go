package admin

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(router *gin.RouterGroup) {
	adminGroup := router.Group("/admin")
	{
		adminGroup.GET("/ping", ping)
	}
}
