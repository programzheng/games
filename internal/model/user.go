package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID       uint
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt sql.NullTime
}

type UserAgent struct {
	ID           uint
	AgentID      uint
	UserID       uint
	ThirdPartyID string
	CreateAt     time.Time
	UpdateAt     time.Time
	DeleteAt     sql.NullTime
}
