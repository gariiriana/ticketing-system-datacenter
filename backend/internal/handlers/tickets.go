package handlers

import (
	"net/http"

	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Save to database
	ticket.ID = "TICKET-001" // Mock ID
	ticket.Status = models.StatusPending

	c.JSON(http.StatusCreated, gin.H{
		"message": "Ticket created successfully",
		"data":    ticket,
	})
}

func UpdateTicketStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status models.TicketStatus `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update in database
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket status updated",
		"id":      id,
		"status":  input.Status,
	})
}
