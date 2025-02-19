package order

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/response"
	"github.com/LordRadamanthys/teste_meli/src/configuration/metrics"
)

func (o *OrderService) FindOrder(idOrder string) (*response.OrdersResponse, error) {

	order, err := o.OrdersOutputPort.FindOrderById(idOrder)
	if err != nil {
		metrics.NotFoundOrders.Inc()
		return nil, err
	}

	processedItens := []response.Item{}
	notProcessedItens := []response.Item{}

	response := response.NewResponse(order, notProcessedItens, processedItens, idOrder)
	return response, nil
}
