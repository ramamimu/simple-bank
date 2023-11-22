package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
}

func (s *Server) Start() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(s.port)
}
