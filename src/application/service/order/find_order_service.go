package order

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/response"
)

func (o *OrderService) FindOrder(idOrder string) (*response.OrdersResponse, error) {

	order, err := o.OrdersOutputPort.FindOrderById(idOrder)
	if err != nil {
		return nil, err
	}

	processedItens := []response.Item{}
	notProcessedItens := []response.Item{}

	response := response.NewResponse(order, notProcessedItens, processedItens, idOrder)
	return response, nil
}
