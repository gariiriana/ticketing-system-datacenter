package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gariiriana/ticketing-system-datacenter/backend/internal/models"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) Create(ctx context.Context, ticket *models.Ticket) error {
	query := `
		INSERT INTO tickets (id, user_id, site_id, description, status, photo_url, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING created_at
	`
	err := r.db.QueryRowContext(ctx, query,
		ticket.ID,
		ticket.UserID,
		ticket.SiteID,
		ticket.Description,
		ticket.Status,
		ticket.PhotoURL,
		time.Now(),
	).Scan(&ticket.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}

	return nil
}

func (r *TicketRepository) GetAll(ctx context.Context) ([]models.Ticket, error) {
	query := `SELECT id, user_id, site_id, description, status, photo_url, created_at, approved_by, approved_at, rejection_reason FROM tickets ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tickets: %w", err)
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var t models.Ticket
		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.SiteID,
			&t.Description,
			&t.Status,
			&t.PhotoURL,
			&t.CreatedAt,
			&t.ApprovedBy,
			&t.ApprovedAt,
			&t.RejectionReason,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ticket: %w", err)
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}

func (r *TicketRepository) UpdateStatus(ctx context.Context, id string, status models.TicketStatus, approvedBy string, reason string) error {
	query := `
		UPDATE tickets 
		SET status = $1, approved_by = $2, approved_at = $3, rejection_reason = $4
		WHERE id = $5
	`
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, status, approvedBy, now, reason, id)
	if err != nil {
		return fmt.Errorf("failed to update ticket status: %w", err)
	}

	return nil
}
