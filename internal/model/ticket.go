package model

import (
	"database/sql"
	"time"
)

type Ticket struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type UserTicket struct {
	ID        uint
	UserID    uint
	TicketID  uint
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
