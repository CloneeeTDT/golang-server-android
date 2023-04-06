package models

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
