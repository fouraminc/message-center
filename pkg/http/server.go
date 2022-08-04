package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type server struct {
	router *gin.Engine
}

func NewServer() *server {
	s := &server{}
	s.init()
	return s
}
func (s *server) init() {
	s.router = gin.Default()
	s.routes()
}

func (s *server) Run() error {
	if err := s.router.Run("localhost:8080"); err != nil {
		return err
	}
	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
