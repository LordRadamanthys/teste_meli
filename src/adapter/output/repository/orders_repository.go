package repository

import (
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/google/uuid"
)

func NewOrderRepository() *Orders {
	return &Orders{}
}

func (o *Orders) SaveOrder(order domain.OrderDomain) string {

	ordertemp := make(map[string]domain.OrderDomain)
	id := uuid.New().String()
	ordertemp[id] = order

	return id
}

func (o *Orders) FindOrderById(orderId string) (*domain.OrderDomain, error) {

	return nil, nil
}
