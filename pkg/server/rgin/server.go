package rgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router *gin.Engine
	Server *http.Server
}

func newServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

// DefaultServer call newServer with default config
func DefaultServer() *Server {
	return newServer()
}

func (s *Server) Start() error {

	logrus.Println("R Gin Server Starting...")

	s.Server = &http.Server{
		Handler:        s.Router,
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.Server.ListenAndServe()
	return nil
}

func (s *Server) Stop() error {
	return s.Server.Close()
}
