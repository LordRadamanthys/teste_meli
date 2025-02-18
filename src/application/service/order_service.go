package service

import (
	"fmt"
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/output"
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

func (o *OrderService) ProcessOrder(order request.OrderRequest) (string, error) {
	var items []domain.ItemDomain
	for _, item := range order.Itens {
		items = append(items, domain.ItemDomain{
			ID: item.ID,
		})
	}
	orderDomain := domain.NewOrderDomain(items)
	var wg sync.WaitGroup

	jobsChan := make(chan string, len(orderDomain.Itens))
	resultChan := make(chan domain.ItemDomain, len(orderDomain.Itens))

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go workerCds(jobsChan, resultChan, o.DistributionCenterOutputPort, &wg)
	}

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

	id := o.OrdersOutputPort.SaveOrder(*orderDomain)

	fmt.Println(orderDomain.Itens)

	return id, nil
}

func (o *OrderService) FindOrder(idOrder string) (*domain.OrderDomain, error) {

	order, err := o.OrdersOutputPort.FindOrderById(idOrder)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func workerCds(jobs <-chan string, result chan<- domain.ItemDomain, dc output.DistributionCenterOutputPort, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range jobs {
		centersList, err := dc.FindDistributionCenterByItemId(item)

		if err != nil {
			result <- domain.ItemDomain{ID: item, DistributionCenter: []string{""}, Processed: false}
		} else {
			result <- domain.ItemDomain{ID: item, DistributionCenter: centersList.AvailableDistributionCenter, Processed: true}
		}

	}
}
