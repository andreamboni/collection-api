package models

import (
	"database/sql"
	"time"
)

type Item struct {
	ID          int64
	Collection  string
	Title       string
	Author      string
	Publisher   string
	ItemType    string
	ItemFormat  string
	PagesNumber int64
	Edition     string
	EditionYear string
	Binding     string
	Language    string
	Country     string
	Copies      int64
}

type ItemResponse struct {
	ID          int64        `json:"id,omitempty"`
	Collection  string       `json:"collection"`
	Title       string       `json:"title"`
	Author      string       `json:"author"`
	Publisher   string       `json:"publisher"`
	ItemType    string       `json:"itemType"`
	ItemFormat  string       `json:"itemFormat"`
	PagesNumber int64        `json:"pagesNumber"`
	Edition     string       `json:"edition"`
	EditionYear string       `json:"editionYear"`
	Binding     string       `json:"binding"`
	Language    string       `json:"language"`
	Country     string       `json:"country"`
	Copies      int64        `json:"copies"`
	CreateDate  time.Time    `json:"createDate"`
	UpdateDate  time.Time    `json:"updateDate"`
	DeleteDate  sql.NullTime `json:"deleteDate,omitempty"`
}
