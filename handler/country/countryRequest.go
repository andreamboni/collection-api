package country

import (
	"fmt"
)

type CreateCountryRequest struct {
	Name string
}

func (r *CreateCountryRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	return nil
}
