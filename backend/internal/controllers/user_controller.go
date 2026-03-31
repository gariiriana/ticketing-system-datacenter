package controllers

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	authClient *auth.Client
}

func NewUserController(authClient *auth.Client) *UserController {
	return &UserController{authClient: authClient}
}

func (ctrl *UserController) SyncUser(c *gin.Context) {
	// 1. Get user claims from context (set by AuthMiddleware)
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims not found"})
		return
	}
	claims := claimsRaw.(map[string]interface{})
	uid := c.GetString("user_id")
	email := ""
	if e, ok := claims["email"].(string); ok {
		email = e
	}

	// 2. Lookup user in DB by email
	var id, role string
	err := repository.DB.QueryRow("SELECT id, role FROM users WHERE email = $1", email).Scan(&id, &role)
	if err != nil {
		// If user not in DB, create as engineer by default
		role = "engineer"
		id = uid // Use Firebase UID as DB ID
		_, err = repository.DB.Exec("INSERT INTO users (id, name, email, role) VALUES ($1, $2, $3, $4)", uid, "User", email, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in DB: " + err.Error()})
			return
		}
	} else if id != uid {
		// Update DB ID to match Firebase UID if its different (mapping existing mock data)
		_, err = repository.DB.Exec("UPDATE users SET id = $1 WHERE email = $2", uid, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync UID to DB: " + err.Error()})
			return
		}
	}

	// 3. Set Custom Claims in Firebase
	err = ctrl.authClient.SetCustomUserClaims(context.Background(), uid, map[string]interface{}{
		"role": role,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set custom claims: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User synced successfully",
		"role":    role,
		"uid":     uid,
	})
}
