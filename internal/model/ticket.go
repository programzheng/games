package model

import (
	"database/sql"
	"time"
)

type Ticket struct {
	ID       uint
	Name     string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt sql.NullTime
}

type UserTicket struct {
	ID        uint
	UserID    uint
	TicketID  uint
	Code      string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt sql.NullTime
}
