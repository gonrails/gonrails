package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/one-hole/gonrails/routers/admin"
	"github.com/one-hole/gonrails/routers/outer"
)

func getRouters() *gin.Engine {
	gin.SetMode(os.Getenv("GO_ENV"))
	router := gin.Default()
	root := router.Group("")
	{
		outer.ApplyRoutes(root)
		admin.ApplyRoutes(root)
	}
	return router
}
