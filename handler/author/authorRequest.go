package author

import (
	"fmt"

	"collection.com/handler"
)

type CreateAuthorRequest struct {
	Name        string
	Nationality string
	// BirthDate   string
	// DeathDate   string
}

func (r *CreateAuthorRequest) Validate() error {
	if r.Name == "" && r.Nationality == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return handler.ErrParamIsRequired("name", "string")
	}
	if r.Nationality == "" {
		return handler.ErrParamIsRequired("nationality", "string")
	}
	// if r.BirthDate == "" {
	// 	return handler.ErrParamIsRequired("birthDate", "string")
	// }
	return nil
}
