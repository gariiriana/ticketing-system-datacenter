package controllers

import (
	"net/http"

	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/models"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type TicketController struct {
	svc *service.TicketService
}

func NewTicketController(svc *service.TicketService) *TicketController {
	return &TicketController{svc: svc}
}

func (ctrl *TicketController) CreateTicket(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.UserID = userID

	if err := ctrl.svc.CreateTicket(c.Request.Context(), &ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

func (ctrl *TicketController) GetTickets(c *gin.Context) {
	// TODO: Filter based on role if needed
	tickets, err := ctrl.svc.GetTickets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (ctrl *TicketController) UpdateStatus(c *gin.Context) {
	userID := c.GetString("user_id")
	claims, _ := c.Get("claims")
	userClaims := claims.(map[string]interface{})
	role, _ := userClaims["role"].(string)

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update ticket status"})
		return
	}

	id := c.Param("id")
	var input struct {
		Status string `json:"status" binding:"required"`
		Reason string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if input.Status == string(models.StatusApproved) {
		err = ctrl.svc.ApproveTicket(c.Request.Context(), id, userID)
	} else if input.Status == string(models.StatusRejected) {
		err = ctrl.svc.RejectTicket(c.Request.Context(), id, userID, input.Reason)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}
