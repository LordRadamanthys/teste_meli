package order

import (
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/output"
	"github.com/LordRadamanthys/teste_meli/src/application/service"
)

type OrderService struct {
	DistributionCenterOutputPort output.DistributionCenterOutputPort
	OrdersOutputPort             output.OrdersOutputPort
}

func NewOrderService(distributionCenterOutputPort output.DistributionCenterOutputPort,
	ordersOutputPort output.OrdersOutputPort) *OrderService {
	return &OrderService{
		DistributionCenterOutputPort: distributionCenterOutputPort,
		OrdersOutputPort:             ordersOutputPort,
	}
}

func (o *OrderService) ProcessOrder(order request.OrderRequest,
	jobsChan chan string, resultChan chan domain.ItemDomain) (string, error) {
	orderDomain := domain.OrderDomain{}

	orderDomain.NewDomain(order)
	var wg sync.WaitGroup

	service.StartWorkers(jobsChan, resultChan, o.DistributionCenterOutputPort, &wg)

	for _, item := range orderDomain.Itens {
		jobsChan <- item.ID
	}

	close(jobsChan)
	wg.Wait()
	close(resultChan)

	var tempItens []domain.ItemDomain
	for item := range resultChan {
		tempItens = append(tempItens, item)
	}

	orderDomain.Itens = tempItens

	id := o.OrdersOutputPort.SaveOrder(orderDomain)

	// fmt.Println(orderDomain.Itens)

	return id, nil
}
