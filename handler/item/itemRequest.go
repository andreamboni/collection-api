package item

import (
	"fmt"

	"collection.com/handler"
)

type CreateItemRequest struct {
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

func (r *CreateItemRequest) Validate() error {
	if r.Title == "" && r.Author == "" && r.Publisher == "" &&
		r.ItemType == "" && r.ItemFormat == "" && r.PagesNumber <= 0 &&
		r.Edition == "" && r.EditionYear == "" && r.Binding == "" &&
		r.Language == "" && r.Country == "" && r.Copies <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Title == "" {
		return handler.ErrParamIsRequired("title", "string")
	}
	if r.Author == "" {
		return handler.ErrParamIsRequired("author", "string")
	}
	if r.Publisher == "" {
		return handler.ErrParamIsRequired("publisher", "string")
	}
	if r.ItemType == "" {
		return handler.ErrParamIsRequired("itemType", "string")
	}
	if r.ItemFormat == "" {
		return handler.ErrParamIsRequired("itemFormat", "string")
	}
	if r.PagesNumber <= 0 {
		return handler.ErrParamIsRequired("pagesNumber", "int64")
	}
	if r.Edition == "" {
		return handler.ErrParamIsRequired("edition", "string")
	}
	if r.EditionYear == "" {
		return handler.ErrParamIsRequired("editionYear", "string")
	}
	if r.Binding == "" {
		return handler.ErrParamIsRequired("binding", "string")
	}
	if r.Language == "" {
		return handler.ErrParamIsRequired("language", "string")
	}
	if r.Country == "" {
		return handler.ErrParamIsRequired("country", "string")
	}
	if r.Copies <= 0 {
		return handler.ErrParamIsRequired("copies", "int64")
	}
	return nil
}
