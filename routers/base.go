package routers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := routers()
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func routers() *gin.Engine {
	gin.SetMode(os.Getenv("GO_ENV"))
	router := gin.Default()
	router.Group("")
	{

	}
	return router
}
