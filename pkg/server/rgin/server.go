package rgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router *gin.Engine
	server *http.Server
}

func newRouter(ginConfig *RouterConfiguration) *gin.Engine {
	gin.SetMode(ginConfig.Mode)
	router := gin.Default()
	if ginConfig.EnableRecovery {
		router.Use(gin.Recovery())
	}
	return router
}

func newServer(config *Configuration) *Server {
	newRouter := newRouter(config.routerConfiguration)
	return &Server{
		Router: newRouter,
		server: &http.Server{
			Handler:        newRouter,
			Addr:           config.httpConfiguration.Addr,
			ReadTimeout:    config.httpConfiguration.ReadTimeout,
			WriteTimeout:   config.httpConfiguration.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

/* -------------------------------- Functions -------------------------------- */
// DefaultServer call newServer with default config
func DefaultServer() *Server {
	logrus.Warnln("RGin Server Initializing By Default Configuration ...")
	return newServer(defaultConfiguration())
}

func NewServer(config *Configuration) *Server {
	return newServer(config)
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
