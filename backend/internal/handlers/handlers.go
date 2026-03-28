package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "up",
		"message": "API is running",
	})
}

// TODO: Implement actual ticket handlers
func GetTickets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of tickets placeholder",
		"data":    []interface{}{},
	})
}
