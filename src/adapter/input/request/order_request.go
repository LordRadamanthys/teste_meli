package request

import "errors"

type OrderRequest struct {
	Items []ItemRequest `json:"items"`
}

func (or *OrderRequest) ValidateRequest() error {
	if len(or.Items) > 100 {
		return errors.New("the number of items cannot exceed 100")
	}

	if len(or.Items) == 0 {
		return errors.New("items cannot be empty")
	}
	return nil
}
