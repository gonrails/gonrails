package rgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (s *Server) Start() error {
	s.Server = &http.Server{
		Handler: s.Router,
	}
	return nil
}

func (s *Server) Stop() error {
	return s.Server.Close()
}
