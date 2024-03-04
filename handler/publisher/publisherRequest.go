package publisher

import (
	"fmt"
)

type CreatePublisherRequest struct {
	Name        string
	Nationality string
}

func (r *CreatePublisherRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	return nil
}
