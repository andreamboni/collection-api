package models

import (
	"database/sql"
	"time"
)

type Author struct {
	ID          int64
	Name        string
	Nationality string
	// BirthDate   time.Time
	// DeathDate   sql.NullTime
}

type AuthorResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	// BirthDate   string       `json:"birthDate"`
	// DeathDate   string       `json:"deathDate,omitempty"`
	CreateDate time.Time    `json:"createDate"`
	UpdateDate time.Time    `json:"updateDate"`
	DeleteDate sql.NullTime `json:"deleteDate,omitempty"`
}
