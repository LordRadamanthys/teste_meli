package order

import "github.com/LordRadamanthys/teste_meli/src/application/domain"

func (o *OrderService) FindOrder(idOrder string) (*domain.OrderDomain, error) {

	order, err := o.OrdersOutputPort.FindOrderById(idOrder)
	if err != nil {
		return nil, err
	}
	return order, nil
}
