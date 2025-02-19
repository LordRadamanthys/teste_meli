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

	processedItems := []response.Item{}
	notProcessedItems := []response.Item{}

	response := response.NewResponse(order, notProcessedItems, processedItems, idOrder)
	return response, nil
}
