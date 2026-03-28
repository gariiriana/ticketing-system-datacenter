package models

import "time"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleEngineer UserRole = "engineer"
)

type User struct {
	ID    string   `json:"id" db:"id"`
	Name  string   `json:"name" db:"name"`
	Email string   `json:"email" db:"email"`
	Role  UserRole `json:"role" db:"role"`
}

type Site struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}

type TicketStatus string

const (
	StatusPending  TicketStatus = "pending"
	StatusApproved TicketStatus = "approved"
	StatusRejected TicketStatus = "rejected"
)

type Ticket struct {
	ID          string       `json:"id" db:"id"`
	UserID      string       `json:"user_id" db:"user_id"`
	SiteID      string       `json:"site_id" db:"site_id"`
	Description string       `json:"description" db:"description"`
	Status      TicketStatus `json:"status" db:"status"`
	PhotoURL    string       `json:"photo_url" db:"photo_url"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	ApprovedBy  string       `json:"approved_by" db:"approved_by"`
	ApprovedAt  *time.Time   `json:"approved_at" db:"approved_at"`
}
