package repository

import (
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/google/uuid"
)

type Orders struct {
	Mu     sync.Mutex
	Orders map[string]domain.OrderDomain `json:"orders"`
}

type OrdersEntity struct {
	ID    string             `json:"id"`
	Order domain.OrderDomain `json:"order"`
}

func (od *Orders) NewOrderRegister(order domain.OrderDomain) string {
	id := uuid.New().String()
	od.Orders[id] = order
	return id
}
