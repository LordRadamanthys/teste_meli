package repository

import (
	"errors"
	"fmt"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/google/uuid"
)

func NewOrderRepository() *Orders {
	return &Orders{
		Order: make(map[string]domain.OrderDomain),
	}
}

func (o *Orders) SaveOrder(order domain.OrderDomain) string {

	id := uuid.New().String()
	o.mu.Lock()
	o.Order[id] = order
	o.mu.Unlock()
	return id
}

func (o *Orders) FindOrderById(orderId string) (*domain.OrderDomain, error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for key, order := range o.Order {
		if key == orderId {
			return &order, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("order with id %s not found", orderId))
}
