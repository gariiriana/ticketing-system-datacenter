package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/config"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/controllers"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/middleware"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/repository"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Initialize Firebase
	opt := option.WithCredentialsFile(cfg.FirebaseCredentialPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting auth client: %v\n", err)
	}

	// Initialize Database
	repository.InitDB(cfg.DBURL)

	// Initialize Repository, Service, and Controller
	ticketRepo := repository.NewTicketRepository(repository.DB)
	ticketSvc := service.NewTicketService(ticketRepo)
	ticketCtrl := controllers.NewTicketController(ticketSvc)

	// Set Gin Mode
	gin.SetMode(cfg.GinMode)

	r := gin.Default()

	// API Routes
	api := r.Group("/api/v1")
	{
		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(authClient))
		{
			api.GET("/tickets", ticketCtrl.GetTickets)
			api.POST("/tickets", ticketCtrl.CreateTicket)
			api.PATCH("/tickets/:id", ticketCtrl.UpdateStatus)
		}
	}

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
