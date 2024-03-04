package models

import (
	"database/sql"
	"time"
)

type Collection struct {
	ID   int64
	Name string
}

type CollectionResponse struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	CreateDate time.Time    `json:"createDate"`
	UpdateDate time.Time    `json:"updateDate"`
	DeleteDate sql.NullTime `json:"deleteDate,omitempty"`
}
