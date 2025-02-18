package repository

import (
	"fmt"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

func NewOrderRepository() *Orders {
	return &Orders{
		Order: make(map[string]domain.OrderDomain),
	}
}

func (o *Orders) SaveOrder(order domain.OrderDomain) string {

	// id := uuid.New().String()
	o.Mu.Lock()
	// o.Order[id] = order
	id := o.NewOrderRegister(order)
	o.Mu.Unlock()
	return id
}

func (o *Orders) FindOrderById(orderId string) (*domain.OrderDomain, error) {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	value, ok := o.Order[orderId]
	if !ok {
		return nil, fmt.Errorf("order with id %s not found", orderId)
	}

	return &value, nil
}
