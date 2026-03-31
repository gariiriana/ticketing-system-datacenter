package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/config"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/controllers"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/middleware"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/repository"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
)

type serviceAccountKey struct {
	Type        string `json:"type"`
	ProjectID   string `json:"project_id"`
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	TokenURI     string `json:"token_uri"`
}

func main() {
	// Logging to stdout
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("=== TICKETING SYSTEM DATACENTER - BACKEND ===")

	// Load Config
	cfg := config.LoadConfig()
	log.Printf("Port: %s | Mode: %s", cfg.Port, cfg.GinMode)

	// Load service account key
	keyBytes, err := os.ReadFile(cfg.FirebaseCredentialPath)
	if err != nil {
		log.Fatalf("Cannot read Firebase credentials: %v", err)
	}

	var sa serviceAccountKey
	if err := json.Unmarshal(keyBytes, &sa); err != nil {
		log.Fatalf("Cannot parse Firebase credentials: %v", err)
	}

	// Use JWT config to create a custom token source
	// This bypasses x509 PKCS8 parsing issues in some Firebase SDK versions
	jwtConfig := &jwt.Config{
		Email:        sa.ClientEmail,
		PrivateKey:   []byte(sa.PrivateKey),
		PrivateKeyID: sa.PrivateKeyID,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/firebase",
		},
		TokenURL: google.JWTTokenURL,
	}
	tokenSource := jwtConfig.TokenSource(context.Background())

	// Initialize Firebase with custom token source
	opt := option.WithTokenSource(tokenSource)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: sa.ProjectID,
	}, opt)
	if err != nil {
		log.Fatalf("Firebase app init failed: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Firebase auth client failed: %v", err)
	}
	log.Println("✅ Firebase initialized")

	// Initialize Database
	repository.InitDB(cfg.DBURL)
	log.Println("✅ Database connected")

	// Initialize layers
	ticketRepo := repository.NewTicketRepository(repository.DB)
	ticketSvc := service.NewTicketService(ticketRepo)
	ticketCtrl := controllers.NewTicketController(ticketSvc)
	userCtrl := controllers.NewUserController(authClient)

	// Set Gin Mode
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check (no auth)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "ticketing-datacenter",
			"version": "1.0.0",
		})
	})

	// Protected API Routes
	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware(authClient))
	{
		api.GET("/user/sync", userCtrl.SyncUser)
		api.GET("/tickets", ticketCtrl.GetTickets)
		api.POST("/tickets", ticketCtrl.CreateTicket)
		api.PATCH("/tickets/:id", ticketCtrl.UpdateStatus)
	}

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
