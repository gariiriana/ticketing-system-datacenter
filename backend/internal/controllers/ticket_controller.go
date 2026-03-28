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
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set UserID from context (TODO: after auth middleware)
	if ticket.UserID == "" {
		ticket.UserID = "engineer-001" // Mock for now
	}

	if err := ctrl.svc.CreateTicket(c.Request.Context(), &ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

func (ctrl *TicketController) GetTickets(c *gin.Context) {
	tickets, err := ctrl.svc.GetTickets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (ctrl *TicketController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adminID := "admin-001" // Mock for now

	var err error
	if input.Status == string(models.StatusApproved) {
		err = ctrl.svc.ApproveTicket(c.Request.Context(), id, adminID)
	} else if input.Status == string(models.StatusRejected) {
		err = ctrl.svc.RejectTicket(c.Request.Context(), id, adminID)
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
