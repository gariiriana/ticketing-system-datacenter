package service

import (
	"context"
	"fmt"

	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/models"
	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/repository"
)

type TicketService struct {
	repo *repository.TicketRepository
}

func NewTicketService(repo *repository.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *models.Ticket) error {
	// Add business validation
	if ticket.Description == "" {
		return fmt.Errorf("description is required")
	}
	if ticket.SiteID == "" {
		return fmt.Errorf("site_id is required")
	}

	// Initialize ticket
	ticket.Status = models.StatusPending

	return s.repo.Create(ctx, ticket)
}

func (s *TicketService) GetTickets(ctx context.Context) ([]models.Ticket, error) {
	return s.repo.GetAll(ctx)
}

func (s *TicketService) ApproveTicket(ctx context.Context, id string, adminID string) error {
	return s.repo.UpdateStatus(ctx, id, models.StatusApproved, adminID, "")
}

func (s *TicketService) RejectTicket(ctx context.Context, id string, adminID string, reason string) error {
	return s.repo.UpdateStatus(ctx, id, models.StatusRejected, adminID, reason)
}
