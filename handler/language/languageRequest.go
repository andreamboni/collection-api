package language

import "collection.com/handler"

type CreateLanguageRequest struct {
	Name string
}

func (r *CreateLanguageRequest) Validate() error {
	if r.Name == "" {
		return handler.ErrParamIsRequired("name", "string")
	}
	return nil
}
