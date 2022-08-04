package model

import (
	"database/sql"
	"time"
)

type Ticket struct {
	ID        uint
	Name      string
	CreateAt  time.Time
	UpdatedAt time.Time
	DeleteAt  sql.NullTime
}

type UserTicket struct {
	ID        uint
	UserID    uint
	TicketID  uint
	Code      string
	CreateAt  time.Time
	UpdatedAt time.Time
	DeleteAt  sql.NullTime
}
