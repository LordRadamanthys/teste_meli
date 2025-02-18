package repository

import (
	"fmt"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

func NewOrderRepository() *Orders {
	return &Orders{
		Orders: make(map[string]domain.OrderDomain),
	}
}

func (o *Orders) SaveOrder(order domain.OrderDomain) string {

	o.Mu.Lock()
	id := o.NewOrderRegister(order)
	o.Mu.Unlock()
	return id
}

func (o *Orders) FindOrderById(orderId string) (*OrdersEntity, error) {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	value, ok := o.Orders[orderId]
	if !ok {
		return nil, fmt.Errorf("order with id %s not found", orderId)
	}

	entity := &OrdersEntity{
		ID:    orderId,
		Order: value,
	}

	return entity, nil
}
