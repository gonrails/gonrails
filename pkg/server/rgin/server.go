package rgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *gin.Engine
	server *http.Server
}

func newRouter(configuration *routerConfiguration) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	return router
}

func newServer(config *Configuration) *Server {
	newRouter := newRouter(config.routerConfiguration)
	return &Server{
		router: newRouter,
		server: &http.Server{
			Handler:        newRouter,
			Addr:           config.httpConfiguration.addr,
			ReadTimeout:    config.httpConfiguration.readTimeout,
			WriteTimeout:   config.httpConfiguration.writeTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

/* -------------------------------- Functions -------------------------------- */
// DefaultServer call newServer with default config
func DefaultServer() *Server {
	logrus.Println("RGin Server Initializing...")

	return newServer(defaultConfiguration())
}

/* -------------------------------- Methods -------------------------------- */
func (s *Server) Start() error {
	logrus.Println("RGin Server Starting...")

	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	logrus.Println("RGin Server Closing...")

	return s.server.Close()
}
