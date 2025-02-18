package repository

import (
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/google/uuid"
)

type Orders struct {
	Mu    sync.Mutex
	Order map[string]domain.OrderDomain `json:"order"`
}

func (od *Orders) NewOrderRegister(order domain.OrderDomain) string {
	id := uuid.New().String()
	od.Order[id] = order
	return id
}
