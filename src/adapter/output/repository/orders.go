package repository

import "github.com/LordRadamanthys/teste_meli/src/application/domain"

type Orders struct {
	Order map[string]domain.OrderDomain `json:"order"`
}
