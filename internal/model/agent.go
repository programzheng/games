package model

import (
	"database/sql"
	"time"
)

type Agent struct {
	ID       uint
	Code     string
	Name     string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt sql.NullTime
}
