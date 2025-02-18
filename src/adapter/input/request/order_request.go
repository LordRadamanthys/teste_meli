package request

import "errors"

type OrderRequest struct {
	Itens []ItemRequest `json:"itens"`
}

func (or *OrderRequest) ValidateRequest() error {
	if len(or.Itens) > 1 {
		return errors.New("the number of items cannot exceed 100")
	}

	if len(or.Itens) == 0 {
		return errors.New("itens cannot be empty")
	}
	return nil
}
