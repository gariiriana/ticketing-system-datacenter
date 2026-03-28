package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Health check
	r.GET("/health", handlers.HealthCheck)

	// API Routes
	api := r.Group("/api/v1")
	{
		// Site Routes
		api.GET("/sites", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List Sites"})
		})

		// Ticket Routes
		api.GET("/tickets", handlers.GetTickets)
		api.POST("/tickets", handlers.CreateTicket)
		api.PATCH("/tickets/:id", handlers.UpdateTicketStatus)
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
