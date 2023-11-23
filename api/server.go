package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	port = ":9090"
)

type Server struct{}

func (s *Server) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", s.handlePing)
	return router
}

func (s *Server) handlePing(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
