package http

import "github.com/gin-gonic/gin"

func (s *server) handlePing() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (s *server) handleRoot() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}
