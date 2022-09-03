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
	UserID    sql.NullInt64
	TicketID  sql.NullInt64
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
