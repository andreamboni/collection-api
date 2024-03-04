package collection

import (
	"fmt"
)

type CreateCollectionRequest struct {
	Name string
}

func (r *CreateCollectionRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	return nil
}
