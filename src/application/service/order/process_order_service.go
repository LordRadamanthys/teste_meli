package order

import (
	"math/rand"
	"sync"
	"time"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/output"
	"github.com/LordRadamanthys/teste_meli/src/application/service"
	"github.com/LordRadamanthys/teste_meli/src/configuration/metrics"
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

	for _, item := range orderDomain.Items {
		jobsChan <- item.ID
	}

	close(jobsChan)
	wg.Wait()
	close(resultChan)

	var tempItems []domain.ItemDomain
	for item := range resultChan {
		item.PrimaryDistributionCenter = ChooseRandomCD(item.DistributionCenter)
		tempItems = append(tempItems, item)
	}

	orderDomain.Items = tempItems

	id := o.OrdersOutputPort.SaveOrder(orderDomain)

	metrics.ItemsTotal.Add(float64(len(tempItems)))
	metrics.OrdersTotal.Inc()
	return id, nil
}

func ChooseRandomCD(availableCDs []string) string {
	if len(availableCDs) == 0 {
		return ""
	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return availableCDs[rd.Intn(len(availableCDs))]
}
