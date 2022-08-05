package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type UserAgent struct {
	ID           uint
	AgentID      uint
	UserID       uint
	ThirdPartyID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}
