package controllers

import (
	"kalista/controllers/admin"
	"kalista/controllers/outer"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Start starts Gin server
func Start() {
	router := getRouters()
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

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